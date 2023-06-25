package router

import "github.com/gin-gonic/gin"

func Initialize() {
	router := gin.Default()

	initializeRouters(router)

	err := router.Run(":8080")
	if err != nil {
		print("Error on startup app")
		return
	}
}
