package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juniorleaoo/learn-go/gopportunities/schemas"
	"net/http"
)

// @BasePath /api/v1
// @Summary Update opening
// @Description Update a job opening
// @Tags Openings
// @Accept json
// @Produces json
// @Param id path int true "Opening identification"
// @Success 200 {object} UpdateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening/:id [put]
func UpdateOpening(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "parameter").Error())
		return
	}

	var request UpdateOpeningRequest
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}

	if err := db.Save(&opening).Error; err != nil {
		logger.Errorf("Error on updating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error update opening on database")
		return
	}
	sendSuccess(ctx, "update-opening", opening)
}
