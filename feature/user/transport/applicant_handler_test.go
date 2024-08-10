package transport

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
)

type MockApplicantUsecase struct {
	mock.Mock
}

func (m *MockApplicantUsecase) CreateApplicant(input dto.ApplicantCreateDTO) error {
	args := m.Called(input)
	return args.Error(0)
}

func (m *MockApplicantUsecase) UpdateApplicant(id int, input dto.ApplicantUpdateDTO) error {
	args := m.Called(id, input)
	return args.Error(0)
}

func (m *MockApplicantUsecase) DeleteApplicant(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockApplicantUsecase) FindApplicantByID(id int) (*dto.ApplicantResponseDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.ApplicantResponseDTO), args.Error(1)
}

func TestCreateApplicant(t *testing.T) {
	mockUsecase := new(MockApplicantUsecase)
	handler := NewApplicantHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/applicant", handler.CreateApplicant)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.ApplicantCreateDTO{
			Email:   "test@example.com",
			Name:    "Johnny",
			Surname: "Hoang",
		}
		mockUsecase.On("CreateApplicant", mockInput).Return(nil)

		body := `{"Email": "test@example.com", "Name": "Johnny", "Surname": "Hoang"}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/applicant", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusCreated, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{"Email": "test@example.com", "Name": "Johnny", "Surname": 15}`
		req, err := http.NewRequest(http.MethodPost, "/api/v1/applicant", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestUpdateApplicant(t *testing.T) {
	mockUsecase := new(MockApplicantUsecase)
	handler := NewApplicantHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.PUT("/api/v1/applicant/:id", handler.UpdateApplicant)

	t.Run("success", func(t *testing.T) {
		mockInput := dto.ApplicantUpdateDTO{
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
		mockUsecase.On("UpdateApplicant", 1, mockInput).Return(nil)

		body := `{
			"RoleID": 1, 
			"DepartmentID": 2, 
			"Email": "test@example.com", 
			"Name": "Tony", 
			"Surname": "Quang", 
			"Gender": "Male", 
			"DOB": "2002-09-20 00:00:00 +0000 UTC",
			"Mobile": "0913895987",
			"CountryID": 2,
			"ResidentCountryID": 7,
		}`
		req, err := http.NewRequest(http.MethodPut, "/api/v1/applicant/1", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		body := `{
			"RoleID": abc, 
			"DepartmentID": 2, 
			"Email": "test@example.com", 
			"Name": "Tony", 
			"Surname": "Quang", 
			"Gender": "Male", 
			"DOB": "2002-09-20 00:00:00 +0000 UTC",
			"Mobile": "0913895987",
			"CountryID": 2,
			"ResidentCountryID": 7,
		}`
		req, err := http.NewRequest(http.MethodPut, "/api/v1/applicant/1", strings.NewReader(body))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestDeleteApplicant(t *testing.T) {
	mockUsecase := new(MockApplicantUsecase)
	handler := NewApplicantHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.DELETE("/api/v1/applicant/:id", handler.DeleteApplicant)

	t.Run("success", func(t *testing.T) {
		mockUsecase.On("DeleteApplicant", 1).Return(nil)

		req, err := http.NewRequest(http.MethodDelete, "/api/v1/applicant/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("bad request", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodDelete, "/api/v1/applicant/abc", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusBadRequest, rr.Code)
	})
}

func TestFindApplicantByID(t *testing.T) {
	mockUsecase := new(MockApplicantUsecase)
	handler := NewApplicantHandler(mockUsecase)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/api/v1/applicant/:id", handler.FindApplicantByID)

	t.Run("success", func(t *testing.T) {
		mockApplicant := &dto.ApplicantResponseDTO{
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
		}
		mockUsecase.On("FindApplicantByID", 1).Return(mockApplicant, nil)

		req, err := http.NewRequest(http.MethodGet, "/api/v1/applicant/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code)
		mockUsecase.AssertExpectations(t)
	})

	t.Run("not found", func(t *testing.T) {
		mockUsecase.On("FindApplicantByID", 1).Return(nil, errors.New("Applicant not found"))

		req, err := http.NewRequest(http.MethodGet, "/api/v1/applicant/1", nil)
		assert.NoError(t, err)

		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusNotFound, rr.Code)
	})
}
