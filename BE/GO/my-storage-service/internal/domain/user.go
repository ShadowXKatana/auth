package domain

import (
	"context"
	"errors"
	"time"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Email        string    `gorm:"column:email;uniqueIndex;size:255;not null" json:"email"`
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:now()" json:"updatedAt"`
}

func (User) TableName() string { return "auth_users" }

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}
