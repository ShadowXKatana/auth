package usecase

import (
	"context"

	"github.com/sos/auth/be/go/my-storage-service/internal/domain"
)

type StorageCreateInput struct {
	UserID string `json:"-"`
	Name   string `json:"name"`
}

type StorageUsecase interface {
	CreateStorage(ctx context.Context, input StorageCreateInput) (domain.Storage, error)
	ListStorages(ctx context.Context, userID string) ([]domain.Storage, error)
	GetStorage(ctx context.Context, id string) (domain.Storage, error)
	DeleteStorage(ctx context.Context, id string) error
}

type storageUsecase struct {
	repo domain.StorageRepository
}

func NewStorageUsecase(repo domain.StorageRepository) StorageUsecase {
	return &storageUsecase{repo: repo}
}

func (uc *storageUsecase) CreateStorage(ctx context.Context, input StorageCreateInput) (domain.Storage, error) {
	if input.Name == "" {
		return domain.Storage{}, ErrInvalidInput
	}

	return uc.repo.Create(ctx, domain.Storage{
		UserID: input.UserID,
		Name:   input.Name,
	})
}

func (uc *storageUsecase) ListStorages(ctx context.Context, userID string) ([]domain.Storage, error) {
	return uc.repo.ListByUserID(ctx, userID)
}

func (uc *storageUsecase) GetStorage(ctx context.Context, id string) (domain.Storage, error) {
	return uc.repo.GetByID(ctx, id)
}

func (uc *storageUsecase) DeleteStorage(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}
