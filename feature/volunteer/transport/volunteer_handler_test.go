package transport

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockVolunteerUsecase struct {
	mock.Mock
}

func (m *MockVolunteerUsecase) CreateVolunteer(input dto.VolunteerCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockVolunteerUsecase) UpdateVolunteer(id int, input dto.VolunteerUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockVolunteerUsecase) DeleteVolunteer(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockVolunteerUsecase) FindVolunteerByID(id int) (*dto.VolunteerResponseDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.VolunteerResponseDTO), args.Error(1)
}

func TestCreateVolunteer(t *testing.T) {
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/volunteer", handler.CreateVolunteer)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.VolunteerCreateDTO{
			UserID:       1,
			DepartmentID: 2,
			Status:       0,
		}
		mockUsecase.On("CreateVolunteer", mockInput).Return(nil)

		body := `{"user_id":1,"department_id":2,"status":"Active"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/volunteer", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"user_id":"abc","department_id":2,"status":"Active"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/volunteer", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestUpdateVolunteer(t *testing.T) {
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/v1/volunteer/:id", handler.UpdateVolunteer)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.VolunteerUpdateDTO{
			DepartmentID: 2,
			Status:       1,
		}
		mockUsecase.On("UpdateVolunteer", 1, mockInput).Return(nil)

		body := `{"department_id":2,"status":"Inactive"}`
		req, err := http.NewRequest(http.MethodPut, "/api/v1/volunteer/1", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"department_id":"abc","status":"Inactive"}`
		req, err := http.NewRequest(http.MethodPut, "/api/v1/volunteer/1", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestDeleteVolunteer(t *testing.T) {
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/api/v1/volunteer/:id", handler.DeleteVolunteer)

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("DeleteVolunteer", 1).Return(nil)

		req, err := http.NewRequest(http.MethodDelete, "/api/v1/volunteer/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/api/v1/volunteer/abc", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestFindVolunteerByID(t *testing.T) {
	mockUsecase := new(MockVolunteerUsecase)
	handler := NewVolunteerHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/v1/volunteer/:id", handler.FindVolunteerByID)

	t.Run("success", func(t *testing.T) {
		mockVolunteer := &dto.VolunteerResponseDTO{
			ID:           1,
			UserID:       1,
			DepartmentID: 2,
			Status:       1,
		}
		mockUsecase.On("FindVolunteerByID", 1).Return(mockVolunteer, nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/volunteer/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUsecase.On("FindVolunteerByID", 1).Return(nil, errors.New("volunteer not found"))

		req, err := http.NewRequest(http.MethodGet, "/api/v1/volunteer/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}