package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juniorleaoo/learn-go/gopportunities/docs"
	"github.com/juniorleaoo/learn-go/gopportunities/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRouters(router *gin.Engine) {
	handler.InitializeHandler()

	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	v1 := router.Group(basePath)

	v1.GET("/opening/:id", handler.ShowOpening)
	v1.POST("/opening", handler.CreateOpening)
	v1.DELETE("/opening/:id", handler.DeleteOpening)
	v1.PUT("/opening/:id", handler.UpdateOpening)
	v1.GET("/openings", handler.ListOpenings)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
