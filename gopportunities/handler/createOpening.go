package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/juniorleaoo/learn-go/gopportunities/schemas"
	"net/http"
)

func CreateOpening(ctx *gin.Context) {
	var request CreateOpeningRequest
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Salary:   request.Salary,
		Remote:   *request.Remote,
		Link:     request.Link,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("Error on creating opening: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
