package transport

import (
	"errors"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/gin-gonic/gin"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/dto"

)

// MockSignInUsecase is a mock implementation of the SignInUsecase interface.
type MockSignInUsecase struct {
	mock.Mock
}

func (m *MockSignInUsecase) SignIn(signInDTO dto.SignInRequestDTO) (*dto.SignInResponseDTO, error) {
	args := m.Called(signInDTO)
	return args.Get(0).(*dto.SignInResponseDTO), args.Error(1)
}

func (m *MockSignInUsecase) SignUp(signUpDTO dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, error) {
	args := m.Called(signUpDTO)
	return args.Get(0).(*dto.SignUpResponseDTO), args.Error(1)
}

func TestSignIn(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	// Set up test data
	signInDTO := dto.SignInRequestDTO{Username: "testuser", Password: "password123"}
	signInResponse := &dto.SignInResponseDTO{
		ID:        1,
		Username:  "testuser",
		Email:     "testuser@example.com",
		Token:     "sample-token",
		CreatedAt: "2024-08-19T00:00:00Z",
		UpdatedAt: "2024-08-19T00:00:00Z",
	}

	mockUsecase.On("SignIn", signInDTO).Return(signInResponse, nil)

	// Create a request to send to the handler
	body, _ := json.Marshal(signInDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.SignInResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, signInResponse, &response)
	mockUsecase.AssertExpectations(t)
}

func TestSignUp(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	// Set up test data
	signUpDTO := dto.SignUpRequestDTO{
		Username: "newuser",
		Password: "password123",
		Email:    "newuser@example.com",
	}
	signUpResponse := &dto.SignUpResponseDTO{
		ID:        2,
		Username:  "newuser",
		Email:     "newuser@example.com",
		CreatedAt: "2024-08-19T00:00:00Z",
		UpdatedAt: "2024-08-19T00:00:00Z",
	}

	mockUsecase.On("SignUp", signUpDTO).Return(signUpResponse, nil)

	// Create a request to send to the handler
	body, _ := json.Marshal(signUpDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusCreated, w.Code)
	var response dto.SignUpResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, signUpResponse, &response)
	mockUsecase.AssertExpectations(t)
}

func TestSignIn_InvalidRequest(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	// Create a request with invalid JSON
	req, _ := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer([]byte("invalid json")))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusBadRequest, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid request data", response["error"])
}

func TestSignUp_UsernameAlreadyExists(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	// Set up test data
	signUpDTO := dto.SignUpRequestDTO{
		Username: "existinguser",
		Password: "password123",
		Email:    "existinguser@example.com",
	}

	mockUsecase.On("SignUp", signUpDTO).Return(nil, errors.New("username already exists"))

	// Create a request to send to the handler
	body, _ := json.Marshal(signUpDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusConflict, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "username already exists", response["error"])
}
