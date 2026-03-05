package postgres

import (
	"context"
	"errors"

	"github.com/sos/auth/be/go/my-storage-service/internal/domain"
	"gorm.io/gorm"
)

type storageRepository struct{ db *gorm.DB }

func NewStorageRepository(db *gorm.DB) domain.StorageRepository {
	return &storageRepository{db: db}
}

func (r *storageRepository) Create(ctx context.Context, storage domain.Storage) (domain.Storage, error) {
	if err := r.db.WithContext(ctx).Create(&storage).Error; err != nil {
		return domain.Storage{}, err
	}

	return storage, nil
}

func (r *storageRepository) ListByUserID(ctx context.Context, userID string) ([]domain.Storage, error) {
	var storages []domain.Storage

	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Order("created_at DESC").Find(&storages).Error; err != nil {
		return nil, err
	}

	return storages, nil
}

func (r *storageRepository) GetByID(ctx context.Context, id string) (domain.Storage, error) {
	var storage domain.Storage

	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&storage).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.Storage{}, domain.ErrStorageNotFound
		}

		return domain.Storage{}, err
	}

	return storage, nil
}

func (r *storageRepository) Delete(ctx context.Context, id string) error {
	result := r.db.WithContext(ctx).Where("id = ?", id).Delete(&domain.Storage{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return domain.ErrStorageNotFound
	}

	return nil
}
