package transport

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

func TestSignInHandler_SignIn_Success(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	signInDTO := dto.SignInRequestDTO{
		Username: "testuser",
		Password: "password",
	}
	mockResponse := &dto.SignInResponseDTO{
		ID:        1,
		Username:  "testuser",
		Email:     "test@example.com",
		Token:     "sample-generated-token",
		CreatedAt: "2024-08-14T00:00:00Z",
		UpdatedAt: "2024-08-14T00:00:00Z",
	}
	mockUsecase.On("SignIn", signInDTO).Return(mockResponse, nil)

	data, _ := json.Marshal(signInDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var response dto.SignInResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, *mockResponse, response)
	mockUsecase.AssertExpectations(t)
}

func TestSignInHandler_SignIn_Failure(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	signInDTO := dto.SignInRequestDTO{
		Username: "testuser",
		Password: "wrongpassword",
	}
	mockUsecase.On("SignIn", signInDTO).Return(nil)

	data, _ := json.Marshal(signInDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-in", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "username or password is incorrect", response["error"])
	mockUsecase.AssertExpectations(t)
}

func TestSignInHandler_SignUp_Success(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	signUpDTO := dto.SignUpRequestDTO{
		Username: "newuser",
		Password: "password",
		Email:    "newuser@example.com",
	}
	mockResponse := &dto.SignUpResponseDTO{
		ID:        2,
		Username:  "newuser",
		Email:     "newuser@example.com",
		CreatedAt: "2024-08-14T00:00:00Z",
		UpdatedAt: "2024-08-14T00:00:00Z",
	}
	mockUsecase.On("SignUp", signUpDTO).Return(mockResponse, nil)

	data, _ := json.Marshal(signUpDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	var response dto.SignUpResponseDTO
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, *mockResponse, response)
	mockUsecase.AssertExpectations(t)
}

func TestSignInHandler_SignUp_Failure(t *testing.T) {
	mockUsecase := new(MockSignInUsecase)
	handler := NewSignInHandler(mockUsecase)
	router := gin.Default()
	handler.RegisterRoutes(router)

	signUpDTO := dto.SignUpRequestDTO{
		Username: "existinguser",
		Password: "password",
		Email:    "existinguser@example.com",
	}
	mockUsecase.On("SignUp", signUpDTO).Return(nil)

	data, _ := json.Marshal(signUpDTO)
	req, _ := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(data))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)
	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "username already exists", response["error"])
	mockUsecase.AssertExpectations(t)
}
