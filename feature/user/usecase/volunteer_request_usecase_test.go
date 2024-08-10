package usecase

import (
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockVolunteerRequestRepository struct {
	mock.Mock
}

func (m *mockVolunteerRequestRepository) CreateVolunteerRequest(volRequest *domain.VolunteerRequest) error {
	args := m.Called(volRequest)
	return args.Error(0)
}

func TestCreateVolunteerRequest(t *testing.T)  {
	mockRepo := new(mockVolunteerRequestRepository)
	usecase := NewVolunteerRequestUsecase(mockRepo)
	
	input := dto.VoluteerRequestCreatingDTO{
		UserID: 1,
		Type:   "verification",
		Status: 0,
	}

	mockRepo.On("CreateApplicantRequest", mock.Anything).Return(nil)

	err := usecase.CreateVolunteerRequest(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}