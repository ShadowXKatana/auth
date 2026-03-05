package postgres

import (
	"context"
	"errors"

	"github.com/sos/auth/be/go/my-storage-service/internal/domain"
	"gorm.io/gorm"
)

type itemRepository struct{ db *gorm.DB }

func NewItemRepository(db *gorm.DB) domain.ItemRepository {
	return &itemRepository{db: db}
}

func (r *itemRepository) Create(ctx context.Context, item domain.Item) (domain.Item, error) {
	if err := r.db.WithContext(ctx).Create(&item).Error; err != nil {
		return domain.Item{}, err
	}

	return item, nil
}

func (r *itemRepository) ListByStorageID(ctx context.Context, storageID string) ([]domain.Item, error) {
	var items []domain.Item

	if err := r.db.WithContext(ctx).Where("storage_id = ?", storageID).Order("created_at DESC").Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (r *itemRepository) GetByID(ctx context.Context, id string) (domain.Item, error) {
	var item domain.Item

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Item{}, domain.ErrItemNotFound
		}

		return domain.Item{}, err
	}

	return item, nil
}

func (r *itemRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Item{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrItemNotFound
	}

	return nil
}

func (r *itemRepository) UpdateTags(ctx context.Context, id string, tags string) (domain.Item, error) {
	var item domain.Item

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Item{}, domain.ErrItemNotFound
		}

		return domain.Item{}, err
	}

	item.Tags = tags
	if err := r.db.WithContext(ctx).Save(&item).Error; err != nil {
		return domain.Item{}, err
	}

	return item, nil
}
