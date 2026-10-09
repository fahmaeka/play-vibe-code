package handler

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/fahmaeka/play-vibe-code/backend/internal/model"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func (h *UserHandler) Register(c echo.Context) error {
	var req RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Format payload tidak valid",
		})
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.ToLower(strings.TrimSpace(req.Email))
	req.Password = strings.TrimSpace(req.Password)

	// Validasi input
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Nama wajib diisi",
		})
	}
	if req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Email wajib diisi",
		})
	}
	if _, err := mail.ParseAddress(req.Email); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Format email tidak valid",
		})
	}
	if len(req.Password) < 6 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Password minimal terdiri dari 6 karakter",
		})
	}

	if h.db == nil {
		return c.JSON(http.StatusServiceUnavailable, map[string]string{
			"error": "Database tidak terhubung",
		})
	}

	// Cek apakah email sudah terdaftar
	var existingUser model.User
	if err := h.db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.JSON(http.StatusConflict, map[string]string{
			"error": "Email sudah terdaftar",
		})
	}

	// Hash password menggunakan bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal memproses password",
		})
	}

	user := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := h.db.Create(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal menyimpan user ke database",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Registrasi user berhasil",
		"user": UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		},
	})
}
