package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type ApplicantRequestUsecaseInterface interface {
	CreateApplicantRequest(request dto.ApplicantRequestCreatingDTO) error
}

type ApplicantRequestUsecase struct {
	RequestRepo storage.ApplicantRequestRepositoryInterface
}

func NewApplicantRequestUsecase(requestRepo storage.ApplicantRequestRepositoryInterface) *ApplicantRequestUsecase {
	return &ApplicantRequestUsecase{RequestRepo: requestRepo}
}

func (u *ApplicantRequestUsecase) CreateApplicantRequest(request dto.ApplicantRequestCreatingDTO) error {
	req := &domain.ApplicantRequestDomain{
		UserID: request.UserID,
		Type:   request.Type,
		Status: request.Status,
	}
	return u.RequestRepo.CreateApplicantRequest(req)
}
