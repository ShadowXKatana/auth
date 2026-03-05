package repository

import (
	"context"
	"fmt"
	"sync"

	"github.com/sos/auth/be/go/init-go-gin/internal/domain"
)

type userRepository struct {
	mu    sync.RWMutex
	store map[string]domain.User
	next  int
}

func NewUserRepository() domain.UserRepository {
	return &userRepository{
		store: make(map[string]domain.User),
		next:  1,
	}
}

func (r *userRepository) Create(_ context.Context, user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	user.ID = fmt.Sprintf("%d", r.next)
	r.next++
	r.store[user.ID] = user

	return user, nil
}

func (r *userRepository) List(_ context.Context) ([]domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	users := make([]domain.User, 0, len(r.store))
	for _, user := range r.store {
		users = append(users, user)
	}

	return users, nil
}
