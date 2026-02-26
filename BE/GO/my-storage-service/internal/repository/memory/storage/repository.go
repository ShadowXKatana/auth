package storage

import (
	"context"
	"fmt"
	"sync"

	domain "github.com/sos/auth/be/go/my-storage-service/internal/domain/storage"
)

type repository struct {
	mu    sync.RWMutex
	store map[string]domain.Storage
	next  int
}

func NewRepository() domain.Repository {
	return &repository{
		store: make(map[string]domain.Storage),
		next:  1,
	}
}

func (r *repository) Create(_ context.Context, storage domain.Storage) (domain.Storage, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	storage.ID = fmt.Sprintf("%d", r.next)
	r.next++
	r.store[storage.ID] = storage

	return storage, nil
}

func (r *repository) List(_ context.Context) ([]domain.Storage, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	storages := make([]domain.Storage, 0, len(r.store))
	for _, storage := range r.store {
		storages = append(storages, storage)
	}

	return storages, nil
}
