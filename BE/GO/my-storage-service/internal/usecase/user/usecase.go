package user

import (
	"context"
	"errors"
	"strings"

	domain "github.com/sos/auth/be/go/my-storage-service/internal/domain/user"
	"github.com/sos/auth/be/go/my-storage-service/internal/security"
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

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type TokenService interface {
	Issue(userID, email string) (string, error)
	Parse(token string) (security.Claims, error)
}

type Usecase interface {
	Register(ctx context.Context, input RegisterInput) (AuthResult, error)
	Login(ctx context.Context, input LoginInput) (AuthResult, error)
	Refresh(ctx context.Context, refreshToken string) (AuthResult, error)
	Me(ctx context.Context, email string) (UserInfo, error)
}

type usecase struct {
	repo                domain.Repository
	hasher              PasswordHasher
	accessTokenService  TokenService
	refreshTokenService TokenService
}

func NewUsecase(repo domain.Repository, hasher PasswordHasher, accessTokenService TokenService, refreshTokenService TokenService) Usecase {
	return &usecase{
		repo:                repo,
		hasher:              hasher,
		accessTokenService:  accessTokenService,
		refreshTokenService: refreshTokenService,
	}
}

func (uc *usecase) Register(ctx context.Context, input RegisterInput) (AuthResult, error) {
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

func (uc *usecase) Login(ctx context.Context, input LoginInput) (AuthResult, error) {
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

func (uc *usecase) Refresh(ctx context.Context, refreshToken string) (AuthResult, error) {
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

func (uc *usecase) Me(ctx context.Context, email string) (UserInfo, error) {
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

func (uc *usecase) issueTokens(user domain.User) (AuthResult, error) {
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
