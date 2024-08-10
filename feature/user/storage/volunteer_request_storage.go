package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
)

type VolunteerRequestRepositoryInterface interface {
	CreateVolunteerRequest(volunteerRequest *domain.VolunteerRequest) error
}

type VolunteerRequestRepository struct {
	DB *gorm.DB
}

func NewVolunteerRequestRepository(db *gorm.DB) *VolunteerRequestRepository {
	return &VolunteerRequestRepository{DB: db}
}

func (r *VolunteerRequestRepository) CreateVolunteerRequest(volunteerRequest *domain.VolunteerRequest) error {
	return r.DB.Create(volunteerRequest).Error
}
