package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juniorleaoo/learn-go/gopportunities/schemas"
	"net/http"
)

func DeleteOpening(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "parameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}

	if err := db.Delete(&opening).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id %s", id))
		return
	}

	sendSuccess(ctx, "delete-opening", opening)
}
