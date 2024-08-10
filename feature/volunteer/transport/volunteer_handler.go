package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/usecase"
	"github.com/gin-gonic/gin"
)

type VolunteerHandler struct {
	VolUsecaseH usecase.VolunteerUsecaseInterface
}

func NewVolunteerHandler(volUsecase usecase.VolunteerUsecaseInterface) *VolunteerHandler {
	return &VolunteerHandler{VolUsecaseH: volUsecase}
}

// CreateVolunteer godoc
// @Summary Create volunteer
// @Description Create volunteer
// @Produce json
// @Tags volunteer
// @Param request body dto.VolunteerCreateDTO true "Create Volunteer Request"
// @Success 201 {string} message "Volunteer created successfully"
// @Router /api/v1/volunteer/ [post]
func (h *VolunteerHandler) CreateVolunteer(c *gin.Context) {
	var input dto.VolunteerCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VolUsecaseH.CreateVolunteer(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

// UpdateVolunteer godoc
// @Summary Update volunteer
// @Description Update volunteer
// @Produce json
// @Tags volunteer
// @Param id path int true "Volunteer ID"
// @Param request body dto.VolunteerUpdateDTO true "Update Volunteer Request"
// @Success 200 {string} message "Volunteer updated successfully"
// @Router /api/v1/volunteer/{id} [put]
func (h *VolunteerHandler) UpdateVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volunteer ID"})
		return
	}

	var input dto.VolunteerUpdateDTO
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VolUsecaseH.UpdateVolunteer(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer updated successfully"})
}

// DeleteVolunteer godoc
// @Summary Delete volunteer
// @Description Delete volunteer
// @Produce json
// @Tags volunteer
// @Param id path int true "Volunteer ID"
// @Success 200 {string} message "Volunteer deleted successfully"
// @Router /api/v1/volunteer/{id} [delete]
func (h *VolunteerHandler) DeleteVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volunteer ID"})
		return
	}

	if err = h.VolUsecaseH.DeleteVolunteer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer deleted successfully"})
}

// FindVolunteerByID godoc
// @Summary Find volunteer by ID
// @Description Find volunteer by ID
// @Produce json
// @Tags volunteer
// @Param id path int true "Volunteer ID"
// @Success 200 {object} domain.Volunteer
// @Router /api/v1/volunteer/{id} [get]
func (h *VolunteerHandler) FindVolunteerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	volunteer, err := h.VolUsecaseH.FindVolunteerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer not found"})
		return
	}

	c.JSON(http.StatusOK, volunteer)
}
