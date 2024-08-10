package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"gorm.io/gorm"
)

type AuthenticationStore interface {
	GetUserByEmail(email string, password string) (*domain.User, string)
	RegisterUser(request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error)
}

type AuthenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *AuthenticationRepository {
	return &AuthenticationRepository{db: db}
}
func (r *AuthenticationRepository) GetUserByEmail(email string, password string) (*domain.User, string) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err.Error()
	}
	if user.Status == 0 {
		return nil, "User is inactive"
	}
	if user.Password != password {
		return nil, "Password is incorrect"
	}
	return &user, ""
}

func (r *AuthenticationRepository) RegisterUser(request *dto.RegisterUserRequest) (*dto.RegisterUserResponse, error) {
	user := domain.User{
		Email:    request.Email,
		Name:     request.Name,
		Password: request.Password,
		Status:   1,
	}

	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}

	response := &dto.RegisterUserResponse{
		Message: "User registered successfully",
	}
	return response, nil
}
