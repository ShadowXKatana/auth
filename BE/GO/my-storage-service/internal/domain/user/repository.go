package user

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type Repository interface {
	Create(ctx context.Context, user User) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
}
