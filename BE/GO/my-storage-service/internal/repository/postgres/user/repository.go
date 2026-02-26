package user

import (
	"context"
	"errors"
	"strings"
	"time"

	domain "github.com/sos/auth/be/go/my-storage-service/internal/domain/user"
	"gorm.io/gorm"
)

type userModel struct {
	ID           string    `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email        string    `gorm:"column:email;uniqueIndex;size:255;not null"`
	PasswordHash string    `gorm:"column:password_hash;not null"`
	CreatedAt    time.Time `gorm:"column:created_at;not null;default:now()"`
	UpdatedAt    time.Time `gorm:"column:updated_at;not null;default:now()"`
}

func (userModel) TableName() string {
	return "auth_users"
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	normalizedEmail := strings.ToLower(strings.TrimSpace(user.Email))
	model := userModel{
		Email:        normalizedEmail,
		PasswordHash: user.PasswordHash,
	}

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.User{}, err
	}

	return domain.User{
		ID:           model.ID,
		Email:        model.Email,
		PasswordHash: model.PasswordHash,
	}, nil
}

func (r *repository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	var model userModel

	err := r.db.WithContext(ctx).Where("email = ?", normalizedEmail).First(&model).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, err
	}

	return domain.User{
		ID:           model.ID,
		Email:        model.Email,
		PasswordHash: model.PasswordHash,
	}, nil
}
