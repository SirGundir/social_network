package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"auth_service/models"
	"auth_service/services"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type MockAuthService struct {
	RegisterFunc func(user *models.User) error
	LoginFunc    func(credentials *models.AuthRequest) (string, error)
}

func (m *MockAuthService) Register(user *models.User) error {
	return m.RegisterFunc(user)
}

func (m *MockAuthService) Login(credentials *models.AuthRequest) (string, error) {
	return m.LoginFunc(credentials)
}

func SetupRouter(authService services.AuthServiceInterface) *gin.Engine {
	r := gin.Default()
	AuthRoutes(r, authService)
	return r
}

func TestRegisterRoute(t *testing.T) {
	authService := &MockAuthService{
		RegisterFunc: func(user *models.User) error {
			return nil
		},
	}
	router := SetupRouter(authService)

	user := models.AuthRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(user)

	req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "User created")
}

func TestLoginRoute(t *testing.T) {
	authService := &MockAuthService{
		LoginFunc: func(credentials *models.AuthRequest) (string, error) {
			return "test-token", nil
		},
	}
	router := SetupRouter(authService)

	credentials := models.AuthRequest{
		Email:    "test@example.com",
		Password: "password123",
	}
	body, _ := json.Marshal(credentials)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "token")
}
