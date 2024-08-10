package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockApplicantRepository struct {
	mock.Mock
}

func (m *MockApplicantRepository) CreateApplicant(applicant *domain.ApplicantDomain) error {
	args := m.Called(applicant)
	return args.Error(0)
}

func (m *MockApplicantRepository) UpdateApplicant(applicant *domain.ApplicantDomain) error {
	args := m.Called(applicant)
	return args.Error(0)
}

func (m *MockApplicantRepository) DeleteApplicant(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockApplicantRepository) FindApplicantByID(id int) (*domain.ApplicantDomain, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.ApplicantDomain), args.Error(1)
}

func TestCreateApplicant(t *testing.T) {
	mockRepo := new(MockApplicantRepository)
	usecase := NewApplicantUsecase(mockRepo)

	input := dto.ApplicantCreateDTO{
		Email:   "test@example.com",
		Name:    "Johnny",
		Surname: "Hoang",
	}

	mockRepo.On("CreateApplicant", mock.Anything).Return(nil)

	err := usecase.CreateApplicant(input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateApplicant(t *testing.T) {
	mockRepo := new(MockApplicantRepository)
	usecase := NewApplicantUsecase(mockRepo)

	input := dto.ApplicantUpdateDTO{
		RoleID:            1,
		DepartmentID:      2,
		Email:             "test@example.com",
		Name:              "Tony",
		Surname:           "Quang",
		Gender:            "Male",
		DOB:               "time.Date(2002, 9, 20, 0, 0, 0, 0, time.UTC)",
		Mobile:            "0913895987",
		CountryID:         2,
		ResidentCountryID: 7,
	}
	applicant := &domain.ApplicantDomain{
		ID:                 1,
		RoleID:             1,
		DepartmentID:       2,
		Email:              "test@example.com",
		Password:           "123456",
		Name:               "Johnny",
		Surname:            "Hoang",
		Gender:             "Male",
		DOB:                time.Date(2002, 9, 20, 0, 0, 0, 0, time.UTC),
		Mobile:             "0912345987",
		CountryID:          3,
		ResidentCountryID:  5,
		Avatar:             "http://linktoavatar.com",
		VerificationStatus: 0,
		Status:             0,
	}

	mockRepo.On("FindApplicantByID", 1).Return(applicant, nil)
	mockRepo.On("UpdateApplicant", applicant).Return(nil)

	err := usecase.UpdateApplicant(1, input)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteApplicant(t *testing.T) {
	mockRepo := new(MockApplicantRepository)
	usecase := NewApplicantUsecase(mockRepo)

	mockRepo.On("DeleteApplicant", 1).Return(nil)

	err := usecase.DeleteApplicant(1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestFindApplicantByID(t *testing.T) {
	mockRepo := new(MockApplicantRepository)
	usecase := NewApplicantUsecase(mockRepo)

	applicant := &domain.ApplicantDomain{
		ID:                1,
		RoleID:            1,
		DepartmentID:      2,
		Email:             "test@example.com",
		Name:              "Johnny",
		Surname:           "Hoang",
		Gender:            "Male",
		DOB:               time.Date(2002, 9, 20, 0, 0, 0, 0, time.UTC),
		Mobile:            "0912345987",
		CountryID:         3,
		ResidentCountryID: 5,
	}

	mockRepo.On("FindApplicantByID", 1).Return(applicant, nil)

	result, err := usecase.FindApplicantByID(1)

	assert.NoError(t, err)
	assert.Equal(t, &dto.ApplicantResponseDTO{
		ID:                1,
		RoleID:            1,
		DepartmentID:      2,
		Email:             "test@example.com",
		Name:              "Johnny",
		Surname:           "Hoang",
		Gender:            "Male",
		DOB:               "2002-09-20 00:00:00 +0000 UTC",
		Mobile:            "0912345987",
		CountryID:         3,
		ResidentCountryID: 5,
	}, result)
	mockRepo.AssertExpectations(t)
}

func TestFindApplicantByID_NotFound(t *testing.T) {
	mockRepo := new(MockApplicantRepository)
	usecase := NewApplicantUsecase(mockRepo)

	mockRepo.On("FindApplicantByID", 1).Return(nil, errors.New("record not found"))

	result, err := usecase.FindApplicantByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
