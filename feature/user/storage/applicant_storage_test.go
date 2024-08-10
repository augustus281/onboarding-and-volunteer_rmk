package storage

import (
	"testing"
	"time"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

func TestCreateApplicant(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewApplicantRepository(&mockDB.DB)

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
	mockDB.On("Create", applicant).Return(&gorm.DB{Error: nil})

	err := repo.CreateApplicant(applicant)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateApplicant(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewApplicantRepository(&mockDB.DB)

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
		Status:             1,
	}
	mockDB.On("Save", applicant).Return(&gorm.DB{Error: nil})

	err := repo.UpdateApplicant(applicant)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteApplicant(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewApplicantRepository(&mockDB.DB)

	mockDB.On("Delete", &domain.ApplicantDomain{}, 1).Return(&gorm.DB{Error: nil})

	err := repo.DeleteApplicant(1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestFindApplicantByID(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewApplicantRepository(&mockDB.DB)

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
	mockDB.On("First", &applicant, 1).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(**domain.ApplicantDomain)
		*arg = applicant
	})

	result, err := repo.FindApplicantByID(1)

	assert.NoError(t, err)
	assert.Equal(t, applicant, result)
	mockDB.AssertExpectations(t)
}

func TestFindApplicantByID_NotFound(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewApplicantRepository(&mockDB.DB)

	mockDB.On("First", mock.Anything, 1).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})

	result, err := repo.FindApplicantByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}
