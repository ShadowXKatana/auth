package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/sos/auth/be/go/my-storage-service/internal/domain"
	"gorm.io/gorm"
)

type userRepository struct{ db *gorm.DB }

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	user.Email = strings.ToLower(strings.TrimSpace(user.Email))

	if err := r.db.WithContext(ctx).Create(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User

	err := r.db.WithContext(ctx).Where("email = ?", strings.ToLower(strings.TrimSpace(email))).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, err
	}

	return user, nil
}
