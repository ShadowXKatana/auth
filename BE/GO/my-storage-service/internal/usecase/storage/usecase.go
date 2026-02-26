package storage

import (
	"context"
	"errors"

	domain "github.com/sos/auth/be/go/my-storage-service/internal/domain/storage"
)

var ErrInvalidInput = errors.New("invalid input")

type CreateInput struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type Usecase interface {
	Create(ctx context.Context, input CreateInput) (domain.Storage, error)
	List(ctx context.Context) ([]domain.Storage, error)
}

type usecase struct {
	repo domain.Repository
}

func NewUsecase(repo domain.Repository) Usecase {
	return &usecase{repo: repo}
}

func (uc *usecase) Create(ctx context.Context, input CreateInput) (domain.Storage, error) {
	if input.Name == "" || input.Path == "" {
		return domain.Storage{}, ErrInvalidInput
	}

	storage := domain.Storage{
		Name: input.Name,
		Path: input.Path,
	}

	return uc.repo.Create(ctx, storage)
}

func (uc *usecase) List(ctx context.Context) ([]domain.Storage, error) {
	return uc.repo.List(ctx)
}
