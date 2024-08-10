package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/domain"
	"gorm.io/gorm"
)

type UserIndentityRepositoryInterface interface {
	CreateUserIdentity(identity *domain.UserIdentity) error
	UpdateUserIdentity(identity *domain.UserIdentity) error
	FindUserIdentityByID(id int) (*domain.UserIdentity, error)
}

type UserIdentityRepository struct {
	DB *gorm.DB
}

func NewUserIdentityRepository(db *gorm.DB) *UserIdentityRepository {
	return &UserIdentityRepository{DB: db}
}

func (r *UserIdentityRepository) CreateUserIdentity(identity *domain.UserIdentity) error {
	return r.DB.Create(identity).Error
}

func (r *UserIdentityRepository) UpdateUserIdentity(identity *domain.UserIdentity) error {
	return r.DB.Save(identity).Error
}

func (r *UserIdentityRepository) FindUserIdentityByID(id int) (*domain.UserIdentity, error) {
	var identity domain.UserIdentity
	if err := r.DB.First(&identity, id).Error; err != nil {
		return nil, err
	}
	return &identity, nil
}
