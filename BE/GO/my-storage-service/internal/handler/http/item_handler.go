package http

import (
	"errors"
	nethttp "net/http"

	"github.com/gin-gonic/gin"
	"github.com/sos/auth/be/go/my-storage-service/internal/usecase"
)

// ── Item Handler ────────────────────────────────────────────────────────────

type ItemHandler struct {
	uc usecase.ItemUsecase
}

func NewItemHandler(uc usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{uc: uc}
}

func (h *ItemHandler) Create(c *gin.Context) {
	var req usecase.ItemCreateInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	req.StorageID = c.Param("id")

	created, err := h.uc.CreateItem(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidInput) {
			c.JSON(nethttp.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		c.JSON(nethttp.StatusInternalServerError, gin.H{"message": "failed to create item"})
		return
	}

	c.JSON(nethttp.StatusCreated, created)
}

func (h *ItemHandler) List(c *gin.Context) {
	items, err := h.uc.ListItems(c.Request.Context(), c.Param("id"))
	if err != nil {
		c.JSON(nethttp.StatusInternalServerError, gin.H{"message": "failed to get items"})
		return
	}

	c.JSON(nethttp.StatusOK, items)
}

func (h *ItemHandler) Delete(c *gin.Context) {
	if err := h.uc.DeleteItem(c.Request.Context(), c.Param("itemId")); err != nil {
		c.JSON(nethttp.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	c.JSON(nethttp.StatusOK, gin.H{"success": true})
}

func (h *ItemHandler) UpdateTags(c *gin.Context) {
	var req usecase.ItemUpdateTagsInput
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(nethttp.StatusBadRequest, gin.H{"message": "invalid request body"})
		return
	}

	updated, err := h.uc.UpdateItemTags(c.Request.Context(), c.Param("itemId"), req.Tags)
	if err != nil {
		c.JSON(nethttp.StatusNotFound, gin.H{"message": "item not found"})
		return
	}

	c.JSON(nethttp.StatusOK, updated)
}
