package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserUsecase is a mock implementation of the UserUsecaseInterface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Login(req dto.LoginUserRequest) (*dto.LoginUserTokenResponse, string) {
	args := m.Called(req)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.LoginUserTokenResponse), args.String(1)
	}
	return nil, args.String(1)
}

func (m *MockUserUsecase) RegisterUser(req dto.RegisterUserRequest) (*dto.RegisterUserResponse, string) {
	args := m.Called(req)
	if args.Get(0) != nil {
		return args.Get(0).(*dto.RegisterUserResponse), args.String(1)
	}
	return nil, args.String(1)
}
func TestAuthenticationHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(MockUserUsecase)
	handler := NewAuthenticationHandler(mockUsecase)

	router := gin.Default()
	router.POST("/api/v1/auth/login", handler.Login)

	t.Run("successful login", func(t *testing.T) {
		loginReq := dto.LoginUserRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		loginResp := &dto.LoginUserTokenResponse{
			Token: "mock-token",
		}
		mockUsecase.On("Login", loginReq).Return(loginResp, "")

		w := httptest.NewRecorder()
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response dto.LoginUserTokenResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, loginResp.Token, response.Token)
	})

	t.Run("login with invalid request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "invalid character 'i' looking for beginning of value", response["error"])
	})

	t.Run("login with incorrect credentials", func(t *testing.T) {
		loginReq := dto.LoginUserRequest{
			Email:    "test@example.com",
			Password: "wrong-password",
		}
		mockUsecase.On("Login", loginReq).Return(nil, "invalid credentials")

		w := httptest.NewRecorder()
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "invalid credentials", response["error"])
	})
}

func TestAuthenticationHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(MockUserUsecase)
	handler := NewAuthenticationHandler(mockUsecase)

	router := gin.Default()
	router.POST("/api/v1/auth/register", handler.Register)

	t.Run("successful registration", func(t *testing.T) {
		registerReq := dto.RegisterUserRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		registerResp := &dto.RegisterUserResponse{

			Message: registerReq.Email,
		}
		mockUsecase.On("RegisterUser", registerReq).Return(registerResp, "")

		w := httptest.NewRecorder()
		body, _ := json.Marshal(registerReq)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var response dto.RegisterUserResponse
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, registerResp.Message, response.Message)
	})

	t.Run("register with invalid request", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer([]byte("invalid body")))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "invalid character 'i' looking for beginning of value", response["error"])
	})

	t.Run("register with existing user", func(t *testing.T) {
		registerReq := dto.RegisterUserRequest{
			Email:    "test@example.com",
			Password: "password",
		}
		mockUsecase.On("RegisterUser", registerReq).Return(nil, "user already exists")

		w := httptest.NewRecorder()
		body, _ := json.Marshal(registerReq)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		var response map[string]string
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "user already exists", response["error"])
	})
}
