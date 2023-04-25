package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type School struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

var schools []School

func main() {
	router := gin.Default()
	router.GET("/schools", listSchools)
	router.POST("/schools", createSchool)
	router.GET("/schools/:id", getSchool)
	router.PUT("/schools/:id", updateSchool)
	router.DELETE("/schools/:id", deleteSchool)
	router.Run(":8080")
}

func listSchools(context *gin.Context) {
	context.JSON(http.StatusOK, schools)
}

func createSchool(context *gin.Context) {
	var school School
	err := context.ShouldBindJSON(&school)
	if err != nil {
		context.AbortWithStatus(http.StatusBadRequest)
		return
	}

	school.ID = len(schools) + 1
	schools = append(schools, school)

	context.JSON(http.StatusCreated, school)
}

func getSchool(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "ID is not int"})
	}

	for _, school := range schools {
		if id == school.ID {
			context.JSON(http.StatusOK, school)
			return
		}
	}

	context.AbortWithStatus(http.StatusNotFound)
}

func updateSchool(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "ID is not int"})
	}

	for i, school := range schools {
		if id == school.ID {
			var updatedSchool School
			err := context.ShouldBindJSON(&updatedSchool)
			if err != nil {
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}

			updatedSchool.ID = school.ID
			schools[i] = updatedSchool

			context.JSON(http.StatusOK, updatedSchool)
			return
		}
	}

	context.AbortWithStatus(http.StatusNotFound)
}

func deleteSchool(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "ID is not int"})
	}

	for i, school := range schools {
		if school.ID == id {
			schools = append(schools[:i], schools[i+1:]...)
			context.Status(http.StatusNoContent)
			return
		}
	}

	context.AbortWithStatus(http.StatusNotFound)
}
