package domain

import (
	"context"
	"errors"
	"time"
)

var ErrItemNotFound = errors.New("item not found")

type Item struct {
	ID        string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"ID"`
	StorageID string    `gorm:"column:storage_id;index;not null" json:"StorageID"`
	Name      string    `gorm:"column:name;size:255;not null" json:"Name"`
	SizeMb    float64   `gorm:"column:size_mb;not null;default:0" json:"SizeMb"`
	Tags      string    `gorm:"column:tags;type:text;not null;default:''" json:"-"`
	CreatedAt time.Time `gorm:"column:created_at;not null;default:now()" json:"CreatedAt"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:now()" json:"UpdatedAt"`
}

func (Item) TableName() string { return "items" }

type ItemRepository interface {
	Create(ctx context.Context, item Item) (Item, error)
	ListByStorageID(ctx context.Context, storageID string) ([]Item, error)
	GetByID(ctx context.Context, id string) (Item, error)
	Delete(ctx context.Context, id string) error
	UpdateTags(ctx context.Context, id string, tags string) (Item, error)
}
