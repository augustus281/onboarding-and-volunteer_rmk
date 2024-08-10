package transport

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockApplicantRequestUsecase struct {
	mock.Mock
}

func (m *MockApplicantRequestUsecase) CreateApplicantRequest(input dto.ApplicantRequestCreatingDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func TestCreateApplicantRequest(t *testing.T) {
	mockUsecase := new(MockApplicantRequestUsecase)
	handler := NewApplicantRequestHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/applicant-request", handler.CreateApplicantRequest)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.ApplicantRequestCreatingDTO{
			UserID: 1,
			Type:   "Application",
			Status: 0,
		}
		mockUsecase.On("CreateApplicantRequest", mockInput).Return(nil)

		body := `{"user_id":1,"type":"Application","status":"Pending"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/applicant-request", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"user_id":abc,"type":"Application","status":"Pending"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/applicant-request", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
