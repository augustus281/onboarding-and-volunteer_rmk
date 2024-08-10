package transport

import (
	"net/http"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
)

type VolunteerRequestHandler struct {
	VolRequestUsecase usecase.VolunteerRequestUsecaseInterface
}

func NewVolunteerRequestHandler(volRequestUsecase usecase.VolunteerRequestUsecaseInterface) *VolunteerRequestHandler {
	return &VolunteerRequestHandler{VolRequestUsecase: volRequestUsecase}
}

// CreateRequest godoc
// @Summary Create a new volunteer request
// @Description Create a new volunteer request
// @Produce json
// @Tags volunteer
// @Accept json
// @Param request body dto.VoluteerRequestCreatingDTO true "Request body"
// @Success 201 {object} string
// @Router /api/v1/volunteer/request [post]
func (h *VolunteerRequestHandler) CreateVolunteerRequest(c *gin.Context) {
	var request dto.VoluteerRequestCreatingDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VolRequestUsecase.CreateVolunteerRequest(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Request created successfully"})
}
