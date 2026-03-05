package http

import (
	"errors"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/my-storage-service/internal/usecase"
)

// ── Storage Handler ─────────────────────────────────────────────────────────

type StorageHandler struct {
	uc usecase.StorageUsecase
}

func NewStorageHandler(uc usecase.StorageUsecase) *StorageHandler {
	return &StorageHandler{uc: uc}
}

func (h *StorageHandler) Create(c *gin.Context) {
	var req usecase.StorageCreateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	userID, _ := c.Get("auth_user_id")
	req.UserID, _ = userID.(string)

	created, err := h.uc.CreateStorage(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			c.JSON(nethttp.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(nethttp.StatusInternalServerError, gin.H{"message": "failed to create storage"})
		return
	}

	c.JSON(nethttp.StatusCreated, created)
}

func (h *StorageHandler) List(c *gin.Context) {
	userID, _ := c.Get("auth_user_id")
	userIDStr, _ := userID.(string)

	storages, err := h.uc.ListStorages(c.Request.Context(), userIDStr)
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"message": "failed to get storages"})
		return
	}

	c.JSON(nethttp.StatusOK, storages)
}

func (h *StorageHandler) GetByID(c *gin.Context) {
	s, err := h.uc.GetStorage(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(nethttp.StatusNotFound, gin.H{"message": "storage not found"})
		return
	}

	c.JSON(nethttp.StatusOK, s)
}

func (h *StorageHandler) Delete(c *gin.Context) {
	if err := h.uc.DeleteStorage(c.Request.Context(), c.Param("id")); err != nil {
		c.JSON(nethttp.StatusNotFound, gin.H{"message": "storage not found"})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"success": true})
}
