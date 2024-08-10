package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
)

type MockDB struct {
	mock.Mock
	gorm.DB
}

func (m *MockDB) Create(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Save(value interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Delete(value interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(value, conds)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) *gorm.DB {
	args := m.Called(dest, conds)
	return args.Get(0).(*gorm.DB)
}

func TestCreateVolunteer(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewVolunteerRepository(&mockDB.DB)
	volunteer := &domain.Volunteer{
		ID:           1,
		UserID:       5,
		DepartmentID: 3,
		Status:       0,
	}
	mockDB.On("Create", volunteer).Return(&gorm.DB{Error: nil})

	err := repo.CreateVolunteer(volunteer)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestUpdateVolunteer(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewVolunteerRepository(&mockDB.DB)

	volunteer := &domain.Volunteer{
		ID:           1,
		DepartmentID: 4,
		Status:       0,
	}
	mockDB.On("Save", volunteer).Return(&gorm.DB{Error: nil})

	err := repo.UpdateVolunteer(volunteer)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestDeleteVolunteer(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewVolunteerRepository(&mockDB.DB)

	mockDB.On("Delete", &domain.Volunteer{}, 1).Return(&gorm.DB{Error: nil})

	err := repo.DeleteVolunteer(1)

	assert.NoError(t, err)
	mockDB.AssertExpectations(t)
}

func TestFindVolunteerByID(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewVolunteerRepository(&mockDB.DB)

	volunteer := &domain.Volunteer{
		ID:           1,
		DepartmentID: 4,
		Status:       0,
	}
	mockDB.On("First", &volunteer, 1).Return(&gorm.DB{Error: nil}).Run(func(args mock.Arguments) {
		arg := args.Get(0).(**domain.Volunteer)
		*arg = volunteer
	})

	result, err := repo.FindVolunteerByID(1)

	assert.NoError(t, err)
	assert.Equal(t, volunteer, result)
	mockDB.AssertExpectations(t)
}

func TestFindVolunteerByID_NotFound(t *testing.T) {
	mockDB := new(MockDB)
	repo := NewVolunteerRepository(&mockDB.DB)

	mockDB.On("First", mock.Anything, 1).Return(&gorm.DB{Error: gorm.ErrRecordNotFound})

	result, err := repo.FindVolunteerByID(1)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockDB.AssertExpectations(t)
}
