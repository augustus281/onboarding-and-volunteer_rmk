package usecase

import (
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/storage"
)

type UserIdentityUsecaseInterface interface {
	CreateUserIdentity(request dto.CreateUserIdentityRequest) error
	UpdateUserIdentity(id int, request dto.UpdateUserIdentityRequest) error
	FindUserIdentityByID(id int) (*dto.UserIdentityResponse, error)
}

type UserIdentityUsecase struct {
	UserIdentityRepo storage.UserIndentityRepositoryInterface
}

func NewUserIdentityUsecase(userIdentityRepo storage.UserIndentityRepositoryInterface) *UserIdentityUsecase {
	return &UserIdentityUsecase{UserIdentityRepo: userIdentityRepo}
}

func (u *UserIdentityUsecase) CreateUserIdentity(request dto.CreateUserIdentityRequest) error {
	expiryDate, err := time.Parse("2006-01-02", request.ExpiryDate)
	if err != nil {
		return err
	}

	identity := &domain.UserIdentity{
		UserID:      request.UserID,
		Number:      request.Number,
		Type:        request.Type,
		Status:      request.Status,
		ExpiryDate:  expiryDate,
		PlaceIssued: request.PlaceIssued,
	}
	return u.UserIdentityRepo.CreateUserIdentity(identity)
}

func (u *UserIdentityUsecase) UpdateUserIdentity(id int, request dto.UpdateUserIdentityRequest) error {
	expiryDate, err := time.Parse("2006-01-02", request.ExpiryDate)
	if err != nil {
		return err
	}
	identity := &domain.UserIdentity{
		ID:          id,
		UserID:      request.UserID,
		Number:      request.Number,
		Type:        request.Type,
		Status:      request.Status,
		ExpiryDate:  expiryDate,
		PlaceIssued: request.PlaceIssued,
	}
	return u.UserIdentityRepo.UpdateUserIdentity(identity)
}

func (u *UserIdentityUsecase) FindUserIdentityByID(id int) (*dto.UserIdentityResponse, error) {

	identity, err := u.UserIdentityRepo.FindUserIdentityByID(id)
	if err != nil {
		return nil, err
	}

	response := &dto.UserIdentityResponse{
		ID:          identity.ID,
		UserID:      identity.UserID,
		Number:      identity.Number,
		Type:        identity.Type,
		Status:      identity.Status,
		ExpiryDate:  identity.ExpiryDate.Format("2006-01-02"),
		PlaceIssued: identity.PlaceIssued,
	}

	return response, nil
}
