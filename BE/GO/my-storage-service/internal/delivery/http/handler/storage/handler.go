package storage

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/my-storage-service/internal/usecase/storage"
)

type Handler struct {
	usecase storage.Usecase
}

func NewHandler(usecase storage.Usecase) *Handler {
	return &Handler{usecase: usecase}
}

func (h *Handler) Create(c *gin.Context) {
	var req storage.CreateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	createdStorage, err := h.usecase.Create(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, storage.ErrInvalidInput) {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to create storage"})
		return
	}

	c.JSON(http.StatusCreated, createdStorage)
}

func (h *Handler) List(c *gin.Context) {
	storages, err := h.usecase.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "failed to get storages"})
		return
	}

	c.JSON(http.StatusOK, storages)
}
