package usecase

import (
	"context"
	"errors"
	"strings"

	"github.com/sos/auth/be/go/my-storage-service/internal/domain"
	"github.com/sos/auth/be/go/my-storage-service/pkg"
)

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrEmailTaken         = errors.New("email already registered")
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type RegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResult struct {
	AccessToken  string   `json:"accessToken"`
	RefreshToken string   `json:"refreshToken"`
	User         UserInfo `json:"user"`
}

type UserInfo struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type UserUsecase interface {
	Register(ctx context.Context, input RegisterInput) (AuthResult, error)
	Login(ctx context.Context, input LoginInput) (AuthResult, error)
	Refresh(ctx context.Context, refreshToken string) (AuthResult, error)
	Me(ctx context.Context, email string) (UserInfo, error)
}

type userUsecase struct {
	repo                domain.UserRepository
	hasher              pkg.PasswordHasher
	accessTokenService  pkg.TokenService
	refreshTokenService pkg.TokenService
}

func NewUserUsecase(repo domain.UserRepository, hasher pkg.PasswordHasher, accessTokenService pkg.TokenService, refreshTokenService pkg.TokenService) UserUsecase {
	return &userUsecase{
		repo:                repo,
		hasher:              hasher,
		accessTokenService:  accessTokenService,
		refreshTokenService: refreshTokenService,
	}
}

func (uc *userUsecase) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
	email := strings.TrimSpace(strings.ToLower(input.Email))
	if !isValidInput(email, input.Password) {
		return AuthResult{}, ErrInvalidInput
	}

	_, err := uc.repo.GetByEmail(ctx, email)
	if err == nil {
		return AuthResult{}, ErrEmailTaken
	}
	if !errors.Is(err, domain.ErrUserNotFound) {
		return AuthResult{}, err
	}

	passwordHash, err := uc.hasher.Hash(input.Password)
	if err != nil {
		return AuthResult{}, err
	}

	createdUser, err := uc.repo.Create(ctx, domain.User{
		Email:        email,
		PasswordHash: passwordHash,
	})
	if err != nil {
		return AuthResult{}, err
	}

	return uc.issueTokens(createdUser)
}

func (uc *userUsecase) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
	email := strings.TrimSpace(strings.ToLower(input.Email))
	if !isValidInput(email, input.Password) {
		return AuthResult{}, ErrInvalidInput
	}

	storedUser, err := uc.repo.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return AuthResult{}, ErrInvalidCredentials
		}

		return AuthResult{}, err
	}

	if err := uc.hasher.Compare(storedUser.PasswordHash, input.Password); err != nil {
		return AuthResult{}, ErrInvalidCredentials
	}

	return uc.issueTokens(storedUser)
}

func (uc *userUsecase) Refresh(ctx context.Context, refreshToken string) (AuthResult, error) {
	if strings.TrimSpace(refreshToken) == "" {
		return AuthResult{}, ErrInvalidCredentials
	}

	claims, err := uc.refreshTokenService.Parse(refreshToken)
	if err != nil {
		return AuthResult{}, ErrInvalidCredentials
	}

	storedUser, err := uc.repo.GetByEmail(ctx, claims.Email)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return AuthResult{}, ErrInvalidCredentials
		}

		return AuthResult{}, err
	}

	return uc.issueTokens(storedUser)
}

func (uc *userUsecase) Me(ctx context.Context, email string) (UserInfo, error) {
	normalizedEmail := strings.TrimSpace(strings.ToLower(email))
	if normalizedEmail == "" {
		return UserInfo{}, ErrInvalidCredentials
	}

	storedUser, err := uc.repo.GetByEmail(ctx, normalizedEmail)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return UserInfo{}, ErrInvalidCredentials
		}

		return UserInfo{}, err
	}

	return UserInfo{ID: storedUser.ID, Email: storedUser.Email}, nil
}

func isValidInput(email, password string) bool {
	if email == "" || password == "" {
		return false
	}

	if len(password) < 6 {
		return false
	}

	return strings.Contains(email, "@")
}

func (uc *userUsecase) issueTokens(user domain.User) (AuthResult, error) {
	accessToken, err := uc.accessTokenService.Issue(user.ID, user.Email)
	if err != nil {
		return AuthResult{}, err
	}

	refreshToken, err := uc.refreshTokenService.Issue(user.ID, user.Email)
	if err != nil {
		return AuthResult{}, err
	}

	return AuthResult{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: UserInfo{
			ID:    user.ID,
			Email: user.Email,
		},
	}, nil
}
