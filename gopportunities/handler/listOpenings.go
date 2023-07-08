package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/juniorleaoo/learn-go/gopportunities/schemas"
	"net/http"
)

func ListOpenings(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
	}

	sendSuccess(ctx, "list-openings", openings)
}
