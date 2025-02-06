package repo

import (
	"context"
	"github.com/minhhoanq/lifeat/user_service/internal/entity"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionRepository interface {
	CreateSession(context.Context, CreateSessionRepoParams) (*entity.Session, error)
	GetSessionByUserId(ctx context.Context, user_id uuid.UUID) (*entity.Session, error)
}

type sessionRepository struct {
	*gorm.DB
}

func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{
		DB: db,
	}
}

type CreateSessionRepoParams struct {
	UserId       uuid.UUID `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiredAt    time.Time `json:"expired_at"`
}

func (s *sessionRepository) CreateSession(ctx context.Context, arg CreateSessionRepoParams) (*entity.Session, error) {
	var session *entity.Session = &entity.Session{
		UserId:       arg.UserId,
		RefreshToken: arg.RefreshToken,
		UserAgent:    arg.UserAgent,
		ClientIp:     arg.ClientIp,
		IsBlocked:    arg.IsBlocked,
		ExpiredAt:    arg.ExpiredAt,
	}

	result := s.DB.WithContext(ctx).Create(&session)

	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}

func (s *sessionRepository) GetSessionByUserId(ctx context.Context, user_id uuid.UUID) (*entity.Session, error) {
	var session *entity.Session
	result := s.DB.WithContext(ctx).First(&session).Where("user_id = ?", user_id)
	if result.Error != nil {
		return nil, result.Error
	}

	return session, nil
}
