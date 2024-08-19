package storage

import (
	"errors"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/sign_in/domain"
	"gorm.io/gorm"
)

type SignInRepositoryImpl struct {
	DB *gorm.DB
}

// NewSignInRepository creates a new instance of SignInRepositoryImpl.
func NewSignInRepository(db *gorm.DB) domain.SignInRepository {
	return &SignInRepositoryImpl{DB: db}
}

// CreateSignIn inserts a new sign-in record into the database.
func (r *SignInRepositoryImpl) CreateSignIn(signIn *domain.SignIn) error {
	result := r.DB.Create(signIn)
	return result.Error
}

// GetSignInByUsername retrieves a sign-in record by username.
func (r *SignInRepositoryImpl) GetSignInByUsername(username string) (*domain.SignIn, error) {
	var signIn domain.SignIn
	result := r.DB.Where("username = ?", username).First(&signIn)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &signIn, result.Error
}

// GetSignInByEmail retrieves a sign-in record by email.
func (r *SignInRepositoryImpl) GetSignInByEmail(email string) (*domain.SignIn, error) {
	var signIn domain.SignIn
	result := r.DB.Where("email = ?", email).First(&signIn)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &signIn, result.Error
}

// UpdateSignIn updates an existing sign-in record in the database.
func (r *SignInRepositoryImpl) UpdateSignIn(signIn *domain.SignIn) error {
	result := r.DB.Save(signIn)
	return result.Error
}

// DeleteSignIn deletes a sign-in record from the database.
func (r *SignInRepositoryImpl) DeleteSignIn(id uint) error {
	var signIn domain.SignIn
	result := r.DB.Delete(&signIn, id)
	return result.Error
}

