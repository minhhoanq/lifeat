package entity

import (
	"time"

	"github.com/google/uuid"
)

type VerifyEmail struct {
	ID         int       `json:"id"`
	UserId     uuid.UUID `json:"user_id"`
	Email      string    `json:"email"`
	SecretCode string    `json:"secret_code"`
	IsUsed     bool      `json:"is_used"`
	ExpiredAt  time.Time `json:"expired_at"`
	CreatedAt  time.Time `json:"created_at"`
}
