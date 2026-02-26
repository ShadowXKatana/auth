package user

import (
	"context"
	"errors"

	domain "github.com/sos/auth/be/go/init-go-gin/internal/domain/user"
)

var ErrInvalidInput = errors.New("invalid input")

type CreateInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Usecase interface {
	Create(ctx context.Context, input CreateInput) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
}

type usecase struct {
	repo domain.Repository
}

func NewUsecase(repo domain.Repository) Usecase {
	return &usecase{repo: repo}
}

func (uc *usecase) Create(ctx context.Context, input CreateInput) (domain.User, error) {
	if input.Name == "" || input.Email == "" {
		return domain.User{}, ErrInvalidInput
	}

	user := domain.User{
		Name:  input.Name,
		Email: input.Email,
	}

	return uc.repo.Create(ctx, user)
}

func (uc *usecase) List(ctx context.Context) ([]domain.User, error) {
	return uc.repo.List(ctx)
}
