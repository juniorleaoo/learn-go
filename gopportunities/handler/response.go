package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sendError(ctx *gin.Context, code int, message string) {
	/*ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message": message,
		"code":    code,
	})*/

	ctx.AbortWithStatusJSON(code, gin.H{
		"message": message,
		"code":    code,
	})
}

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", operation),
		"data":    data,
	})
}
