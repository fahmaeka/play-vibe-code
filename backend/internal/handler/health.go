package handler

import (
	"net/http"
	"time"

	"github.com/fahmaeka/play-vibe-code/backend/internal/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var startTime = time.Now()

type HealthHandler struct {
	cfg *config.Config
	db  *gorm.DB
}

func NewHealthHandler(cfg *config.Config, db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		cfg: cfg,
		db:  db,
	}
}

func (h *HealthHandler) Check(c echo.Context) error {
	dbStatus := "disconnected"
	if h.db != nil {
		sqlDB, err := h.db.DB()
		if err == nil {
			if err := sqlDB.Ping(); err == nil {
				dbStatus = "connected"
			}
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      "ok",
		"service":     "play-vibe-backend",
		"environment": h.cfg.AppEnv,
		"uptime":      time.Since(startTime).String(),
		"database":    dbStatus,
		"timestamp":   time.Now().Format(time.RFC3339),
	})
}
