package user

import "context"

type Repository interface {
	Create(ctx context.Context, user User) (User, error)
	List(ctx context.Context) ([]User, error)
}
