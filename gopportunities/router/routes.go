package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juniorleaoo/learn-go/gopportunities/handler"
)

func initializeRouters(router *gin.Engine) {
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")

	v1.GET("/opening/:id", handler.ShowOpening)
	v1.POST("/opening", handler.CreateOpening)
	v1.DELETE("/opening/:id", handler.DeleteOpening)
	v1.PUT("/opening/:id", handler.UpdateOpening)
	v1.GET("/openings", handler.ListOpenings)

}
