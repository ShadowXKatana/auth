package usecase

import (
	"context"
	"errors"

	"github.com/sos/auth/be/go/init-go-gin/internal/domain"
)

var ErrInvalidInput = errors.New("invalid input")

type CreateUserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserUsecase interface {
	Create(ctx context.Context, input CreateUserInput) (domain.User, error)
	List(ctx context.Context) ([]domain.User, error)
}

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (uc *userUsecase) Create(ctx context.Context, input CreateUserInput) (domain.User, error) {
	if input.Name == "" || input.Email == "" {
		return domain.User{}, ErrInvalidInput
	}

	user := domain.User{
		Name:  input.Name,
		Email: input.Email,
	}

	return uc.repo.Create(ctx, user)
}

func (uc *userUsecase) List(ctx context.Context) ([]domain.User, error) {
	return uc.repo.List(ctx)
}
