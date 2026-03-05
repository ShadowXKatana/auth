package http

import (
	"errors"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/init-go-gin/internal/usecase"
)

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase: usecase}
}

func (h *UserHandler) Create(c *gin.Context) {
	var req usecase.CreateUserInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	createdUser, err := h.usecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			c.JSON(nethttp.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(nethttp.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	c.JSON(nethttp.StatusCreated, createdUser)
}

func (h *UserHandler) List(c *gin.Context) {
	users, err := h.usecase.List(c.Request.Context())
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"message": "failed to get users"})
		return
	}

	c.JSON(nethttp.StatusOK, users)
}
