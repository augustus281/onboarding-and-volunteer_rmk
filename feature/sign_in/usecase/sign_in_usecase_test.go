package usecase

import (
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

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

func TestSignIn_Success(t *testing.T) {
	repo := new(MockSignInRepository)
	usecase := NewSignInUsecase(repo)

	// Prepare mock data
	mockSignIn := &domain.SignIn{
		ID:           1,
		Username:     "testuser",
		PasswordHash: hashPassword("password"), // Assume hashPassword is a helper function
		Email:        "test@example.com",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	repo.On("GetSignInByUsername", "testuser").Return(mockSignIn, nil)

	// Define input DTO
	signInDTO := dto.SignInRequestDTO{
		Username: "testuser",
		Password: "password",
	}

	// Execute SignIn
	response, err := usecase.SignIn(signInDTO)

	// Assert results
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, mockSignIn.ID, response.ID)
	assert.Equal(t, mockSignIn.Username, response.Username)
	assert.Equal(t, mockSignIn.Email, response.Email)
	assert.Equal(t, "sample-generated-token", response.Token)
	repo.AssertExpectations(t)
}

func TestSignIn_Failure(t *testing.T) {
	repo := new(MockSignInRepository)
	usecase := NewSignInUsecase(repo)

	// Setup mock
	repo.On("GetSignInByUsername", "testuser").Return(nil, nil)

	// Define input DTO
	signInDTO := dto.SignInRequestDTO{
		Username: "testuser",
		Password: "password",
	}

	// Execute SignIn
	response, err := usecase.SignIn(signInDTO)

	// Assert results
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "username or password is incorrect", err.Error())
	repo.AssertExpectations(t)
}

func TestSignUp_Success(t *testing.T) {
	repo := new(MockSignInRepository)
	usecase := NewSignInUsecase(repo)

	// Prepare mock data
	repo.On("GetSignInByUsername", "newuser").Return(nil, nil)
	repo.On("GetSignInByEmail", "newuser@example.com").Return(nil, nil)

	// Define input DTO
	signUpDTO := dto.SignUpRequestDTO{
		Username: "newuser",
		Password: "password",
		Email:    "newuser@example.com",
	}

	// Setup expectations
	repo.On("CreateSignIn", mock.Anything).Return(nil)

	// Execute SignUp
	response, err := usecase.SignUp(signUpDTO)

	// Assert results
	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, "newuser", response.Username)
	assert.Equal(t, "newuser@example.com", response.Email)
	repo.AssertExpectations(t)
}

func TestSignUp_Failure_UsernameExists(t *testing.T) {
	repo := new(MockSignInRepository)
	usecase := NewSignInUsecase(repo)

	// Prepare mock data
	existingSignIn := &domain.SignIn{
		Username: "existinguser",
	}
	repo.On("GetSignInByUsername", "newuser").Return(existingSignIn, nil)
	repo.On("GetSignInByEmail", "newuser@example.com").Return(nil, nil)

	// Define input DTO
	signUpDTO := dto.SignUpRequestDTO{
		Username: "newuser",
		Password: "password",
		Email:    "newuser@example.com",
	}

	// Execute SignUp
	response, err := usecase.SignUp(signUpDTO)

	// Assert results
	assert.Error(t, err)
	assert.Nil(t, response)
	assert.Equal(t, "username already exists", err.Error())
	repo.AssertExpectations(t)
}

func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}
