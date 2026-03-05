package http

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/my-storage-service/internal/usecase"
)

// ── Auth Handler ────────────────────────────────────────────────────────────

type AuthHandler struct {
	uc     usecase.UserUsecase
	secure bool
}

func NewAuthHandler(uc usecase.UserUsecase, secure bool) *AuthHandler {
	return &AuthHandler{uc: uc, secure: secure}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req usecase.RegisterInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	result, err := h.uc.Register(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidInput):
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		case errors.Is(err, usecase.ErrEmailTaken):
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to register user"})
		}
		return
	}

	h.setAuthCookies(c, result)
	c.JSON(http.StatusCreated, result)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req usecase.LoginInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	result, err := h.uc.Login(c.Request.Context(), req)
	if err != nil {
		switch {
		case errors.Is(err, usecase.ErrInvalidInput):
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		case errors.Is(err, usecase.ErrInvalidCredentials):
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to login"})
		}
		return
	}

	h.setAuthCookies(c, result)
	c.JSON(http.StatusOK, result)
}

func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil || refreshToken == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "missing refresh token"})
		return
	}

	result, err := h.uc.Refresh(c.Request.Context(), refreshToken)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to refresh token"})
		return
	}

	h.setAuthCookies(c, result)
	c.JSON(http.StatusOK, result)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	h.clearAuthCookies(c)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

func (h *AuthHandler) Me(c *gin.Context) {
	email, _ := c.Get("auth_user_email")
	emailValue, _ := email.(string)

	me, err := h.uc.Me(c.Request.Context(), emailValue)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to load profile"})
		return
	}

	c.JSON(http.StatusOK, me)
}

func (h *AuthHandler) setAuthCookies(c *gin.Context, result usecase.AuthResult) {
	setCookie(c, "access_token", result.AccessToken, getEnvInt("JWT_EXPIRES_IN_MINUTES", 60)*60, h.secure)
	setCookie(c, "refresh_token", result.RefreshToken, getEnvInt("JWT_REFRESH_EXPIRES_IN_MINUTES", 10080)*60, h.secure)
}

func (h *AuthHandler) clearAuthCookies(c *gin.Context) {
	setCookie(c, "access_token", "", -1, h.secure)
	setCookie(c, "refresh_token", "", -1, h.secure)
}

func setCookie(c *gin.Context, name, value string, maxAge int, secure bool) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   maxAge,
		Expires:  time.Now().Add(time.Duration(maxAge) * time.Second),
	})
}

func getEnvInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil || parsed <= 0 {
		return fallback
	}

	return parsed
}
