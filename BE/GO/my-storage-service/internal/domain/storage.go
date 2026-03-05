package domain

import (
	"context"
	"errors"
	"time"
)

var ErrStorageNotFound = errors.New("storage not found")

type Storage struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	UserID    string    `gorm:"column:user_id;index;not null" json:"userId"`
	Name      string    `gorm:"column:name;size:255;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:now()" json:"updatedAt"`
}

func (Storage) TableName() string { return "storages" }

type StorageRepository interface {
	Create(ctx context.Context, storage Storage) (Storage, error)
	ListByUserID(ctx context.Context, userID string) ([]Storage, error)
	GetByID(ctx context.Context, id string) (Storage, error)
	Delete(ctx context.Context, id string) error
}
