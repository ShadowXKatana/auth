package user

import (
	"context"
	"fmt"
	"strings"
	"sync"

	domain "github.com/sos/auth/be/go/my-storage-service/internal/domain/user"
)

type repository struct {
	mu           sync.RWMutex
	usersByID    map[string]domain.User
	usersByEmail map[string]string
	next         int
}

func NewRepository() domain.Repository {
	return &repository{
		usersByID:    make(map[string]domain.User),
		usersByEmail: make(map[string]string),
		next:         1,
	}
}

func (r *repository) Create(_ context.Context, user domain.User) (domain.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	email := strings.ToLower(strings.TrimSpace(user.Email))
	user.ID = fmt.Sprintf("%d", r.next)
	user.Email = email
	r.next++

	r.usersByID[user.ID] = user
	r.usersByEmail[email] = user.ID

	return user, nil
}

func (r *repository) GetByEmail(_ context.Context, email string) (domain.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	userID, ok := r.usersByEmail[normalizedEmail]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}

	storedUser, ok := r.usersByID[userID]
	if !ok {
		return domain.User{}, domain.ErrUserNotFound
	}

	return storedUser, nil
}
