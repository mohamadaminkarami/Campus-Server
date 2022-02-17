package controllers

import (
	"backend/models"
	"backend/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllSchoolCourses(c *gin.Context) {
	var schools []models.School
	models.GetDB().Find(&schools)
	var response []map[string]interface{}

	for _, u := range schools {
		response = append(response, serializers.SchoolCoursesToJSON(u))
	}
	c.JSON(http.StatusOK, gin.H{"result": response})
}

func GetAllSchools(c *gin.Context) {
	var schools []models.School
	models.GetDB().Find(&schools)
	var response []map[string]interface{}

	for _, u := range schools {
		response = append(response, serializers.SchoolToJSON(u))
	}
	c.JSON(http.StatusOK, gin.H{"result": response})
}
