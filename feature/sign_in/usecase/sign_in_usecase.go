package usecase

import (
	"errors"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/dto"
	"golang.org/x/crypto/bcrypt"
)

type SignInUsecase interface {
	SignIn(signInDTO dto.SignInRequestDTO) (*dto.SignInResponseDTO, error)
	SignUp(signUpDTO dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, error)
}

type SignInUsecaseImpl struct {
	signInRepo domain.SignInRepository
}

// NewSignInUsecase creates a new instance of SignInUsecaseImpl.
func NewSignInUsecase(repo domain.SignInRepository) SignInUsecase {
	return &SignInUsecaseImpl{signInRepo: repo}
}

// SignIn handles the business logic for user sign-in.
func (u *SignInUsecaseImpl) SignIn(signInDTO dto.SignInRequestDTO) (*dto.SignInResponseDTO, error) {
	// Find user by username
	signIn, err := u.signInRepo.GetSignInByUsername(signInDTO.Username)
	if err != nil {
		return nil, err
	}
	if signIn == nil {
		return nil, errors.New("username or password is incorrect")
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(signIn.Password), []byte(signInDTO.Password)); err != nil {
		return nil, errors.New("username or password is incorrect")
	}

	// Generate a token (for simplicity, token generation is not shown here)
	token := "sample-generated-token" // Replace with actual token generation logic

	// Return response DTO
	response := &dto.SignInResponseDTO{
		ID:        signIn.ID,
		Username:  signIn.Username,
		Email:     signIn.Email,
		Token:     token,
		CreatedAt: signIn.CreatedAt.Format(time.RFC3339),
		UpdatedAt: signIn.UpdatedAt.Format(time.RFC3339),
	}
	return response, nil
}

// SignUp handles the business logic for user registration.
func (u *SignInUsecaseImpl) SignUp(signUpDTO dto.SignUpRequestDTO) (*dto.SignUpResponseDTO, error) {
	// Check if the username or email already exists
	existingUser, err := u.signInRepo.GetSignInByUsername(signUpDTO.Username)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("username already exists")
	}

	existingEmail, err := u.signInRepo.GetSignInByEmail(signUpDTO.Email)
	if err != nil {
		return nil, err
	}
	if existingEmail != nil {
		return nil, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signUpDTO.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	// Create the SignIn entity
	newSignIn := &domain.SignIn{
		Username:  signUpDTO.Username,
		Password:  string(hashedPassword),
		Email:     signUpDTO.Email,
		UserID:    signUpDTO.UserID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Save the new user
	if err := u.signInRepo.CreateSignIn(newSignIn); err != nil {
		return nil, errors.New("failed to create user")
	}

	// Return response DTO
	response := &dto.SignUpResponseDTO{
		ID:        newSignIn.ID,
		Username:  newSignIn.Username,
		Email:     newSignIn.Email,
		UserID:    newSignIn.UserID,
		CreatedAt: newSignIn.CreatedAt.Format(time.RFC3339),
		UpdatedAt: newSignIn.UpdatedAt.Format(time.RFC3339),
	}
	return response, nil
}
