package usecase

import (
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockApplicantRequestRepository struct {
	mock.Mock
}

func (m *mockApplicantRequestRepository) CreateApplicantRequest(appliRequest *domain.ApplicantRequestDomain) error {
	args := m.Called(appliRequest)
	return args.Error(0)
}

func TestCreateApplicantRequest(t *testing.T) {
	mockRepo := new(mockApplicantRequestRepository)
	usecase := NewApplicantRequestUsecase(mockRepo)
	
	input := dto.ApplicantRequestCreatingDTO{
		UserID: 1,
		Type:   "application",
		Status: 0,
	}

	mockRepo.On("CreateApplicantRequest", mock.Anything).Return(nil)

	err := usecase.CreateApplicantRequest(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
