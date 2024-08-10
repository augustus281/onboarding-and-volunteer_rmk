package storage

import (
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreateVolunteerRequest(t *testing.T)  {
	mockDB := new(MockDB)
	repo := NewVolunteerRequestRepository(&mockDB.DB)
	volRequest := &domain.VolunteerRequest{
		ID:     1,
		UserID: 2,
		Type:   "Application",
		Status: 0,
	}
	mockDB.On("Create", volRequest).Return(&gorm.DB{Error: nil})

	err := repo.CreateVolunteerRequest(volRequest)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}