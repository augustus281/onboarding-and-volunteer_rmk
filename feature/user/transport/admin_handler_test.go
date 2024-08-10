package transport

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAdminUsecase is a mock implementation of the AdminUsecaseInterface
type MockAdminUsecase struct {
	mock.Mock
}

func (m *MockAdminUsecase) GetListPendingRequest() (*dto.ListRequest, string) {
	args := m.Called()
	return args.Get(0).(*dto.ListRequest), args.String(1)
}

func (m *MockAdminUsecase) GetPendingRequestById(id int) (*dto.RequestResponse, string) {
	args := m.Called(id)
	return args.Get(0).(*dto.RequestResponse), args.String(1)
}

func (m *MockAdminUsecase) GetListRequest() (*dto.ListRequest, string) {
	args := m.Called()
	return args.Get(0).(*dto.ListRequest), args.String(1)
}

func (m *MockAdminUsecase) GetRequestById(id int) (*dto.RequestResponse, string) {
	args := m.Called(id)
	return args.Get(0).(*dto.RequestResponse), args.String(1)
}

func (m *MockAdminUsecase) ApproveRequest(id, userId int) string {
	args := m.Called(id, userId)
	return args.String(0)
}

func (m *MockAdminUsecase) RejectRequest(id, userId int) string {
	args := m.Called(id, userId)
	return args.String(0)
}

func (m *MockAdminUsecase) AddRejectNotes(id int, notes string) string {
	args := m.Called(id, notes)
	return args.String(0)
}

func (m *MockAdminUsecase) DeleteRequest(id int) string {
	args := m.Called(id)
	return args.String(0)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.Default()
}

func TestGetListPendingRequest(t *testing.T) {
	mockUsecase := new(MockAdminUsecase)
	handler := NewAuthenticationHandler(mockUsecase)

	router := setupRouter()
	router.GET("/api/v1/admin/list-pending-request", handler.GetListPendingRequest)

	mockUsecase.On("GetListPendingRequest").Return([]dto.RequestResponse{}, "")

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/list-pending-request", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "[]")
	mockUsecase.AssertExpectations(t)
}

func TestGetPendingRequestById(t *testing.T) {
	mockUsecase := new(MockAdminUsecase)
	handler := NewAuthenticationHandler(mockUsecase)

	router := setupRouter()
	router.GET("/api/v1/admin/pending-request/:id", handler.GetPendingRequestById)

	mockUsecase.On("GetPendingRequestById", 1).Return(dto.RequestResponse{}, "")

	req, _ := http.NewRequest(http.MethodGet, "/api/v1/admin/pending-request/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "{}")
	mockUsecase.AssertExpectations(t)
}

// More tests for other handlers can be added similarly...

func TestApproveRequest(t *testing.T) {
	mockUsecase := new(MockAdminUsecase)
	handler := NewAuthenticationHandler(mockUsecase)

	router := setupRouter()
	router.POST("/api/v1/admin/approve-request/:id", func(c *gin.Context) {
		c.Set("userId", 1)
		handler.ApproveRequest(c)
	})

	mockUsecase.On("ApproveRequest", 1, 1).Return("Request approved")

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/approve-request/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Request approved")
	mockUsecase.AssertExpectations(t)
}

func TestRejectRequest(t *testing.T) {
	mockUsecase := new(MockAdminUsecase)
	handler := NewAuthenticationHandler(mockUsecase)

	router := setupRouter()
	router.POST("/api/v1/admin/reject-request/:id", func(c *gin.Context) {
		c.Set("userId", 1)
		handler.RejectRequest(c)
	})

	mockUsecase.On("RejectRequest", 1, 1).Return("Request rejected")

	req, _ := http.NewRequest(http.MethodPost, "/api/v1/admin/reject-request/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Request rejected")
	mockUsecase.AssertExpectations(t)
}
