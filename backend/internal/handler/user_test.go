package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestUserHandler_Register_Validation(t *testing.T) {
	e := echo.New()
	h := NewUserHandler(nil)

	tests := []struct {
		name         string
		payload      string
		expectedCode int
	}{
		{
			name:         "empty name",
			payload:      `{"name":"","email":"test@example.com","password":"secretpassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "empty email",
			payload:      `{"name":"Test User","email":"","password":"secretpassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid email format",
			payload:      `{"name":"Test User","email":"invalid-email","password":"secretpassword"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "short password",
			payload:      `{"name":"Test User","email":"test@example.com","password":"123"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "database unavailable",
			payload:      `{"name":"Test User","email":"test@example.com","password":"secretpassword"}`,
			expectedCode: http.StatusServiceUnavailable,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/api/users/register", strings.NewReader(tt.payload))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			err := h.Register(c)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if rec.Code != tt.expectedCode {
				t.Errorf("expected status %d, got %d. Body: %s", tt.expectedCode, rec.Code, rec.Body.String())
			}
		})
	}
}
