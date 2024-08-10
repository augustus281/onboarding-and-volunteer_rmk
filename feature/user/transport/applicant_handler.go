package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"

	"github.com/gin-gonic/gin"
)

type ApplicantHandler struct {
	ApplicantUseCaseH usecase.ApplicantUsecaseInterface
}

func NewApplicantHandler(userUsecase usecase.ApplicantUsecaseInterface) *ApplicantHandler {
	return &ApplicantHandler{ApplicantUseCaseH: userUsecase}
}

// CreateApplicant godoc
// @Summary Create applicant
// @Description Create applicant
// @Produce json
// @Tags applicant
// @Param request body dto.ApplicantCreateDTO true "Create Applicant Request"
// @Success 201 {string} message "Applicant created successfully"
// @Router /api/v1/applicant/ [post]
func (h *ApplicantHandler) CreateApplicant(c *gin.Context) {
	var request dto.ApplicantCreateDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ApplicantUseCaseH.CreateApplicant(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// UpdateApplicant godoc
// @Summary Update applicant
// @Description Update applicant
// @Produce json
// @Tags applicant
// @Param id path int true "Applicant ID"
// @Param request body dto.AppplicantUpdateDTO true "Update Applicant Request"
// @Success 200 {string} message "Applicant updated successfully"
// @Router /api/v1/applicant/{id} [put]
func (h *ApplicantHandler) UpdateApplicant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var request dto.ApplicantUpdateDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.ApplicantUseCaseH.UpdateApplicant(id, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteApplicant godoc
// @Summary Delete applicant
// @Description Delete applicant
// @Produce json
// @Tags applicant
// @Param id path int true "Applicant ID"
// @Success 200 string message
// @Router /api/v1/applicant/{id} [delete]
func (h *ApplicantHandler) DeleteApplicant(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.ApplicantUseCaseH.DeleteApplicant(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

// FindApplicantByID godoc
// @Summary Find applicant by ID
// @Description Find applicant by ID
// @Produce json
// @Tags applicant
// @Param id path int true "Applicant ID"
// @Success 200 {object} dto.ApplicantResponseDTO
// @Router /api/v1/applicant/{id} [get]
func (h *ApplicantHandler) FindApplicantByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := h.ApplicantUseCaseH.FindApplicantByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
