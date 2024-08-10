package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"

	"gorm.io/gorm"
)

type ApplicantRepositoryInterface interface {
	CreateApplicant(user *domain.ApplicantDomain) error
	UpdateApplicant(user *domain.ApplicantDomain) error
	DeleteApplicant(id int) error
	FindApplicantByID(id int) (*domain.ApplicantDomain, error)
}

type ApplicantRepository struct {
	DB *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) *ApplicantRepository {
	return &ApplicantRepository{DB: db}
}

func (r *ApplicantRepository) CreateApplicant(user *domain.ApplicantDomain) error {
	return r.DB.Create(user).Error
}

func (r *ApplicantRepository) UpdateApplicant(user *domain.ApplicantDomain) error {
	return r.DB.Save(user).Error
}

func (r *ApplicantRepository) DeleteApplicant(id int) error {
	return r.DB.Delete(&domain.ApplicantDomain{}, id).Error
}

func (r *ApplicantRepository) FindApplicantByID(id int) (*domain.ApplicantDomain, error) {
	var user domain.ApplicantDomain
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
