package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/minhhoanq/lifeat/user_service/config"
	"github.com/minhhoanq/lifeat/user_service/internal/entity"
	"github.com/minhhoanq/lifeat/user_service/internal/token"
	"github.com/minhhoanq/lifeat/user_service/internal/usecase/repo"
	"github.com/minhhoanq/lifeat/user_service/internal/util"
	"time"

	"github.com/google/uuid"
)

type UserUsecase interface {
	GetUserByID(context.Context, uuid.UUID) (*entity.User, error)
	CreateUser(context.Context, CreateUserUsecaseParams) (*entity.User, error)
	Login(context.Context, LoginUsecaseParams) (*LoginUscaseResponse, error)
	RenewAccessToken(ctx context.Context, arg RenewAccessTokenUsecaseParams) (*renewAccessTokenUsecaseResponse, error)
}

type userUsecase struct {
	userRepo    repo.UserRepository
	sessionRepo repo.SessionRepository
	tokenMaker  token.Maker
	cfg         config.Config
}

func New(userRepo repo.UserRepository, sessionRepo repo.SessionRepository, tokenMaker token.Maker, cfg config.Config) UserUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		sessionRepo: sessionRepo,
		tokenMaker:  tokenMaker,
		cfg:         cfg,
	}
}

func (uc userUsecase) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := uc.userRepo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type CreateUserUsecaseParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
}

func (uc userUsecase) CreateUser(ctx context.Context, arg CreateUserUsecaseParams) (*entity.User, error) {
	hashedPassword, err := util.HashPassword(arg.Password)
	if err != nil {
		return nil, err
	}

	payload := repo.CreateUserRepoParams{
		Username: arg.Username,
		Email:    arg.Email,
		Password: hashedPassword,
		RoleId:   arg.RoleId,
	}

	user, err := uc.userRepo.CreateUser(ctx, payload)
	if err != nil {
		return nil, err
	}

	return user, nil
}

type LoginUsecaseParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUscaseResponse struct {
	SessionID             int          `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	RefreshToken          string       `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  *entity.User `json:"user"`
}

func (uc userUsecase) Login(ctx context.Context, arg LoginUsecaseParams) (*LoginUscaseResponse, error) {
	user, err := uc.userRepo.GetUserByUsername(ctx, arg.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, err
	}

	// password of user is hashed password
	err = util.ComparePassword(arg.Password, user.Password)
	if err != nil {
		return nil, err
	}

	accessToken, accessPayload, err := uc.tokenMaker.CreateToken(user.ID, user.RoleId, uc.cfg.AccessTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("create access token error: %v", err)
	}

	refreshToken, refreshPayload, err := uc.tokenMaker.CreateToken(user.ID, user.RoleId, uc.cfg.RefreshTokenDuration)
	if err != nil {
		return nil, fmt.Errorf("create refresh token error: %v", err)
	}

	session, err := uc.sessionRepo.CreateSession(ctx, repo.CreateSessionRepoParams{
		UserId:       user.ID,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiredAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, fmt.Errorf("faild create session: %v", err)
	}

	return &LoginUscaseResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  user,
	}, nil
}

type RenewAccessTokenUsecaseParams struct {
	RefreshToken string `json:"refresh_token"`
}

type renewAccessTokenUsecaseResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (uc userUsecase) RenewAccessToken(ctx context.Context, arg RenewAccessTokenUsecaseParams) (*renewAccessTokenUsecaseResponse, error) {
	refreshTokenPayload, err := uc.tokenMaker.VerifyToken(arg.RefreshToken)
	if err != nil {
		return nil, err
	}

	session, err := uc.sessionRepo.GetSessionByUserId(ctx, refreshTokenPayload.UserId)
	if err != nil {
		return nil, err
	}

	if session.IsBlocked {
		return nil, fmt.Errorf("token is blocked")
	}

	if session.UserId != refreshTokenPayload.UserId {
		return nil, fmt.Errorf("incorrect session user")
	}

	if session.RefreshToken != arg.RefreshToken {
		return nil, fmt.Errorf("mismatch session token")
	}

	if time.Now().After(session.ExpiredAt) {
		return nil, fmt.Errorf("expired session")
	}

	user, err := uc.userRepo.GetUserByID(ctx, session.UserId)
	if err != nil {
		return nil, err
	}

	accessToken, payload, err := uc.tokenMaker.CreateToken(user.ID, user.RoleId, uc.cfg.AccessTokenDuration)
	if err != nil {
		return nil, err
	}

	rsp := &renewAccessTokenUsecaseResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: payload.ExpiredAt,
	}

	return rsp, nil
}
