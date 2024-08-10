package storage

import (
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateApplicantRequest(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewApplicantRequestRepository(&mockDB.DB)
	appliRequest := &domain.ApplicantRequestDomain{
		ID:     1,
		UserID: 2,
		Type:   "Application",
		Status: 0,
	}
	mockDB.On("Create", appliRequest).Return(&gorm.DB{Error: nil})

	err := repo.CreateApplicantRequest(appliRequest)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}
