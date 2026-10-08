package main

import (
	"log"

	"github.com/fahmaeka/play-vibe-code/backend/internal/config"
	"github.com/fahmaeka/play-vibe-code/backend/internal/database"
	"github.com/fahmaeka/play-vibe-code/backend/internal/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	// Initialize Database (soft-fail if MySQL isn't up yet, to allow server to still serve /health)
	var db *gorm.DB
	var err error
	db, err = database.InitDB(cfg)
	if err != nil {
		log.Printf("⚠️ Warning: Database connection failed: %v", err)
		log.Println("Server will start in degraded mode (Health endpoint will report DB disconnected).")
	} else {
		log.Println("✅ Successfully connected to MySQL database!")
	}

	e := echo.New()
	e.HideBanner = true

	// Middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.OPTIONS},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Handlers
	healthHandler := handler.NewHealthHandler(cfg, db)
	itemHandler := handler.NewItemHandler(db)

	// Routes
	e.GET("/health", healthHandler.Check)
	e.GET("/", healthHandler.Check)

	api := e.Group("/api")
	{
		api.GET("/items", itemHandler.GetAll)
		api.POST("/items", itemHandler.Create)
		api.DELETE("/items/:id", itemHandler.Delete)
	}

	serverAddr := ":" + cfg.Port
	log.Printf("🚀 Server running on http://localhost%s", serverAddr)
	if err := e.Start(serverAddr); err != nil {
		log.Fatalf("Server stopped: %v", err)
	}
}
