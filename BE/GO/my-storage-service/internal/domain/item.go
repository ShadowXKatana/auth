package domain

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
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

// MarshalJSON serializes Tags as a []string instead of a raw comma-separated string.
func (i Item) MarshalJSON() ([]byte, error) {
	type Alias struct {
		ID        string    `json:"ID"`
		StorageID string    `json:"StorageID"`
		Name      string    `json:"Name"`
		SizeMb    float64   `json:"SizeMb"`
		Tags      []string  `json:"Tags"`
		CreatedAt time.Time `json:"CreatedAt"`
		UpdatedAt time.Time `json:"UpdatedAt"`
	}

	tags := []string{}
	if i.Tags != "" {
		for _, t := range strings.Split(i.Tags, ",") {
			if trimmed := strings.TrimSpace(t); trimmed != "" {
				tags = append(tags, trimmed)
			}
		}
	}

	return json.Marshal(Alias{
		ID:        i.ID,
		StorageID: i.StorageID,
		Name:      i.Name,
		SizeMb:    i.SizeMb,
		Tags:      tags,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
	})
}

type ItemRepository interface {
	Create(ctx context.Context, item Item) (Item, error)
	ListByStorageID(ctx context.Context, storageID string) ([]Item, error)
	GetByID(ctx context.Context, id string) (Item, error)
	Delete(ctx context.Context, id string) error
	UpdateTags(ctx context.Context, id string, tags string) (Item, error)
}
