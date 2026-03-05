package pkg

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type TokenService interface {
	Issue(userID, email string) (string, error)
	Parse(token string) (Claims, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hashedPassword, password string) error
}

type jwtService struct {
	secretKey []byte
	ttl       time.Duration
}

func NewJWTService(secret string, ttl time.Duration) TokenService {
	return &jwtService{
		secretKey: []byte(secret),
		ttl:       ttl,
	}
}

func (s *jwtService) Issue(userID, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.ttl)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *jwtService) Parse(token string) (Claims, error) {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}

		return s.secretKey, nil
	})
	if err != nil {
		return Claims{}, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return Claims{}, ErrInvalidToken
	}

	return *claims, nil
}
