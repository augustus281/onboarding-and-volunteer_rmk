package usecase

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
)

// Mocking the AdminRepositoryInterface
type MockAdminRepository struct {
	mock.Mock
}

func (m *MockAdminRepository) GetListPendingRequest() ([]*domain.Request, string) {
	args := m.Called()
	return args.Get(0).([]*domain.Request), args.String(1)
}

func (m *MockAdminRepository) GetPendingRequestByID(id int) (*domain.Request, string) {
	args := m.Called(id)
	return args.Get(0).(*domain.Request), args.String(1)
}

func (m *MockAdminRepository) GetListAllRequest() ([]*domain.Request, string) {
	args := m.Called()
	return args.Get(0).([]*domain.Request), args.String(1)
}

func (m *MockAdminRepository) GetRequestByID(id int) (*domain.Request, string) {
	args := m.Called(id)
	return args.Get(0).(*domain.Request), args.String(1)
}

func (m *MockAdminRepository) ApproveRequest(id int, verifier_id int) string {
	args := m.Called(id, verifier_id)
	return args.String(0)
}

func (m *MockAdminRepository) RejectRequest(id int, verifier_id int) string {
	args := m.Called(id, verifier_id)
	return args.String(0)
}

func (m *MockAdminRepository) AddRejectNotes(id int, notes string) string {
	args := m.Called(id, notes)
	return args.String(0)
}

func (m *MockAdminRepository) DeleteRequest(id int) string {
	args := m.Called(id)
	return args.String(0)
}

func TestGetListPendingRequest(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	usecase := NewAdminUsecase(mockRepo)
	mockRepo.On("GetListPendingRequest").Return(nil, "No request found")

	result, msg := usecase.GetListPendingRequest()
	assert.Nil(t, result)
	assert.Equal(t, "No request found", msg)
	mockRepo.AssertExpectations(t)
}

func TestGetPendingRequestById(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	usecase := NewAdminUsecase(mockRepo)

	mockRequest := &domain.Request{
		ID:          1,
		UserID:      23,
		Type:        "abc",
		Status:      123,
		RejectNotes: "abc",
		VerifierID:  124,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	mockRepo.On("GetPendingRequestByID", 1).Return(mockRequest, "Request found")

	result, msg := usecase.GetPendingRequestById(1)
	assert.NotNil(t, result)
	assert.Equal(t, "Request found", msg)
	assert.Equal(t, mockRequest.ID, result.ID)
	assert.Equal(t, mockRequest.UserID, result.UserID)
	assert.Equal(t, mockRequest.Type, result.Type)
	assert.Equal(t, mockRequest.Status, result.Status)
	assert.Equal(t, mockRequest.RejectNotes, result.RejectNotes)
	assert.Equal(t, mockRequest.VerifierID, result.VerifierID)
	mockRepo.AssertExpectations(t)
}

func TestApproveRequest(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	usecase := NewAdminUsecase(mockRepo)

	mockRepo.On("ApproveRequest", 1, 456).Return("Request approved")

	msg := usecase.ApproveRequest(1, 456)
	assert.Equal(t, "Request approved", msg)
	mockRepo.AssertExpectations(t)
}

func TestRejectRequest(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	usecase := NewAdminUsecase(mockRepo)

	mockRepo.On("RejectRequest", 1, 456).Return("Request rejected")

	msg := usecase.RejectRequest(1, 456)
	assert.Equal(t, "Request rejected", msg)
	mockRepo.AssertExpectations(t)
}

func TestAddRejectNotes(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	usecase := NewAdminUsecase(mockRepo)

	mockRepo.On("AddRejectNotes", 1, "Some notes").Return("Reject notes added")

	msg := usecase.AddRejectNotes(1, "Some notes")
	assert.Equal(t, "Reject notes added", msg)
	mockRepo.AssertExpectations(t)
}

func TestDeleteRequest(t *testing.T) {
	mockRepo := new(MockAdminRepository)
	usecase := NewAdminUsecase(mockRepo)

	mockRepo.On("DeleteRequest", 1).Return("Request deleted")

	msg := usecase.DeleteRequest(1)
	assert.Equal(t, "Request deleted", msg)
	mockRepo.AssertExpectations(t)
}
