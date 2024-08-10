package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)
type VolunteerRequestUsecaseInterface interface {
	CreateVolunteerRequest(request dto.VoluteerRequestCreatingDTO) error
}

type VolunteerRequestUsecase struct {
	VolRequestRepo storage.VolunteerRequestRepositoryInterface
}

func NewVolunteerRequestUsecase(volRequestRepo storage.VolunteerRequestRepositoryInterface) *VolunteerRequestUsecase {
	return &VolunteerRequestUsecase{VolRequestRepo: volRequestRepo}
}

func (u *VolunteerRequestUsecase) CreateVolunteerRequest(request dto.VoluteerRequestCreatingDTO) error {
	req := &domain.VolunteerRequest{
		UserID: request.UserID,
		Type:   request.Type,
		Status: request.Status,
	}
	return u.VolRequestRepo.CreateVolunteerRequest(req)
}
