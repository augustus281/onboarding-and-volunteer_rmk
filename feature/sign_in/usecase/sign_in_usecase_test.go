package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/dto"
	"golang.org/x/crypto/bcrypt"
)

// MockSignInRepository is a mock implementation of the SignInRepository interface.
type MockSignInRepository struct {
	mock.Mock
}

func (m *MockSignInRepository) CreateSignIn(signIn *domain.SignIn) error {
	args := m.Called(signIn)
	return args.Error(0)
}

func (m *MockSignInRepository) GetSignInByUsername(username string) (*domain.SignIn, error) {
	args := m.Called(username)
	return args.Get(0).(*domain.SignIn), args.Error(1)
}

func (m *MockSignInRepository) GetSignInByEmail(email string) (*domain.SignIn, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.SignIn), args.Error(1)
}

func (m *MockSignInRepository) UpdateSignIn(signIn *domain.SignIn) error {
	args := m.Called(signIn)
	return args.Error(0)
}

func (m *MockSignInRepository) DeleteSignIn(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestSignIn(t *testing.T) {
	mockRepo := new(MockSignInRepository)
	usecase := NewSignInUsecase(mockRepo)

	// Set up test data
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	signIn := &domain.SignIn{
		ID:        1,
		Username:  "testuser",
		Password:  string(hashedPassword),
		Email:     "testuser@example.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetSignInByUsername", "testuser").Return(signIn, nil)

	// Test SignIn
	dto := dto.SignInRequestDTO{Username: "testuser", Password: "password123"}
	response, err := usecase.SignIn(dto)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), response.ID)
	assert.Equal(t, "testuser", response.Username)
	assert.Equal(t, "testuser@example.com", response.Email)
	mockRepo.AssertExpectations(t)
}

func TestSignUp(t *testing.T) {
	mockRepo := new(MockSignInRepository)
	usecase := NewSignInUsecase(mockRepo)

	// Set up test data
	signUpDTO := dto.SignUpRequestDTO{
		Username: "newuser",
		Password: "password123",
		Email:    "newuser@example.com",
	}

	// Set up mock expectations
	mockRepo.On("GetSignInByUsername", "newuser").Return(nil, nil)
	mockRepo.On("GetSignInByEmail", "newuser@example.com").Return(nil, nil)
	mockRepo.On("CreateSignIn", mock.Anything).Return(nil)

	// Test SignUp
	response, err := usecase.SignUp(signUpDTO)
	assert.NoError(t, err)
	assert.Equal(t, "newuser", response.Username)
	assert.Equal(t, "newuser@example.com", response.Email)
	mockRepo.AssertExpectations(t)
}
