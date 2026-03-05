package domain

import "context"

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRepository interface {
	Create(ctx context.Context, user User) (User, error)
	List(ctx context.Context) ([]User, error)
}
