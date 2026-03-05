package usecase

import (
	"context"
	"strings"

	"github.com/sos/auth/be/go/my-storage-service/internal/domain"
)

type ItemCreateInput struct {
	StorageID string   `json:"-"`
	Name      string   `json:"name"`
	SizeMb    float64  `json:"sizeMb"`
	Tags      []string `json:"tags"`
}

type ItemUpdateTagsInput struct {
	Tags []string `json:"tags"`
}

type ItemUsecase interface {
	CreateItem(ctx context.Context, input ItemCreateInput) (domain.Item, error)
	ListItems(ctx context.Context, storageID string) ([]domain.Item, error)
	DeleteItem(ctx context.Context, id string) error
	UpdateItemTags(ctx context.Context, id string, tags []string) (domain.Item, error)
}

type itemUsecase struct {
	repo domain.ItemRepository
}

func NewItemUsecase(repo domain.ItemRepository) ItemUsecase {
	return &itemUsecase{repo: repo}
}

func (uc *itemUsecase) CreateItem(ctx context.Context, input ItemCreateInput) (domain.Item, error) {
	if input.Name == "" || input.StorageID == "" {
		return domain.Item{}, ErrInvalidInput
	}

	return uc.repo.Create(ctx, domain.Item{
		StorageID: input.StorageID,
		Name:      input.Name,
		SizeMb:    input.SizeMb,
		Tags:      joinTags(input.Tags),
	})
}

func (uc *itemUsecase) ListItems(ctx context.Context, storageID string) ([]domain.Item, error) {
	return uc.repo.ListByStorageID(ctx, storageID)
}

func (uc *itemUsecase) DeleteItem(ctx context.Context, id string) error {
	return uc.repo.Delete(ctx, id)
}

func (uc *itemUsecase) UpdateItemTags(ctx context.Context, id string, tags []string) (domain.Item, error) {
	return uc.repo.UpdateTags(ctx, id, joinTags(tags))
}

func joinTags(tags []string) string {
	if len(tags) == 0 {
		return ""
	}
	return strings.Join(tags, ",")
}
