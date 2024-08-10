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

type MockVolunteerRequestUsecase struct {
	mock.Mock
}

func (m *MockVolunteerRequestUsecase) CreateVolunteerRequest(input dto.VoluteerRequestCreatingDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func TestCreateVolunteerRequest(t *testing.T)  {
	mockUsecase := new(MockVolunteerRequestUsecase)
	handler := NewVolunteerRequestHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/volunteer-request", handler.CreateVolunteerRequest)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.VoluteerRequestCreatingDTO{
			UserID: 1,
			Type:   "Verification",
			Status: 0,
		}
		mockUsecase.On("CreateVolunteerRequest", mockInput).Return(nil)

		body := `{"user_id":1,"type":"Verification","status":"Pending"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/volunteer-request", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"user_id":abc,"type":"Verification","status":"Pending"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/volunteer-request", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}