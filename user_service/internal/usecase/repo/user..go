package repo

import (
	"context"
	"fmt"
	"github.com/minhhoanq/lifeat/user_service/internal/entity"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(context.Context, uuid.UUID) (*entity.User, error)
	CreateUser(ctx context.Context, arg CreateUserRepoParams) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
}

type userRepository struct {
	*gorm.DB
}

// New user repository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

// GetUserByID -.
func (u *userRepository) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	fmt.Println("[INFO] UserRepo - GetUser - User ID", id)
	var user *entity.User
	result := u.DB.WithContext(ctx).First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// GetUserByUsername -.
func (u *userRepository) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	fmt.Println("[INFO] UserRepo - GetUser - UserName", username)
	var user *entity.User
	result := u.DB.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

type CreateUserRepoParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	RoleId   int    `json:"role_id"`
}

// CreateUser -.
func (u *userRepository) CreateUser(ctx context.Context, arg CreateUserRepoParams) (*entity.User, error) {
	fmt.Println("[INFO] UserRepo - CreateUser - User", arg)
	var user *entity.User = &entity.User{
		Username: arg.Username,
		Email:    arg.Email,
		Password: arg.Password,
		RoleId:   arg.RoleId,
	}

	result := u.DB.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
