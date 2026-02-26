package storage

import "context"

type Repository interface {
	Create(ctx context.Context, storage Storage) (Storage, error)
	List(ctx context.Context) ([]Storage, error)
}
