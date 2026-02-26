package user

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/init-go-gin/internal/usecase/user"
)

type Handler struct {
	usecase user.Usecase
}

func NewHandler(usecase user.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Create(c *gin.Context) {
	var req user.CreateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	createdUser, err := h.usecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, user.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (h *Handler) List(c *gin.Context) {
	users, err := h.usecase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
