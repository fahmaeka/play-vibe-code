package handler

import (
	"net/http"
	"strconv"

	"github.com/fahmaeka/play-vibe-code/backend/internal/model"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ItemHandler struct {
	db *gorm.DB
}

func NewItemHandler(db *gorm.DB) *ItemHandler {
	return &ItemHandler{db: db}
}

func (h *ItemHandler) GetAll(c echo.Context) error {
	if h.db == nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Database not connected"})
	}

	var items []model.Item
	if err := h.db.Order("id desc").Find(&items).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) Create(c echo.Context) error {
	if h.db == nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Database not connected"})
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
	}

	if req.Title == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Title is required"})
	}

	if req.Status == "" {
		req.Status = "pending"
	}

	item := model.Item{
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
	}

	if err := h.db.Create(&item).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) Delete(c echo.Context) error {
	if h.db == nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{"error": "Database not connected"})
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid item ID"})
	}

	if err := h.db.Delete(&model.Item{}, id).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Item deleted successfully"})
}
