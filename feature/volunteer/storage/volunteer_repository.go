package storage

import (
	"gorm.io/gorm"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
)

// VolunteerRepositoryInterface defines the methods that a VolunteerRepository should implement
type VolunteerRepositoryInterface interface {
	CreateVolunteer(volunteer *domain.Volunteer) error
	UpdateVolunteer(volunteer *domain.Volunteer) error
	DeleteVolunteer(id int) error
	FindVolunteerByID(id int) (*domain.Volunteer, error)
}

type VolunteerRepository struct {
	DB *gorm.DB
}

func NewVolunteerRepository(db *gorm.DB) *VolunteerRepository {
	return &VolunteerRepository{DB: db}
}

func (r *VolunteerRepository) CreateVolunteer(volunteer *domain.Volunteer) error {
	return r.DB.Create(volunteer).Error
}

func (r *VolunteerRepository) UpdateVolunteer(volunteer *domain.Volunteer) error {
	return r.DB.Save(volunteer).Error
}

func (r *VolunteerRepository) DeleteVolunteer(id int) error {
	return r.DB.Delete(&domain.Volunteer{}, id).Error
}

func (r *VolunteerRepository) FindVolunteerByID(id int) (*domain.Volunteer, error) {
	var volunteer *domain.Volunteer
	if err := r.DB.First(&volunteer, id).Error; err != nil {
		return nil, err
	}
	return volunteer, nil
}
