package controllers

import (
	"backend/models"
	"backend/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSchool(c *gin.Context) {
	var School models.School
	if err := c.ShouldBindJSON(&School); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	models.GetDB().Create(&School)
	c.JSON(http.StatusOK, serializers.SchoolToJSON(School))
}

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

func UpdateSchool(c *gin.Context) {
	var newSchool models.School
	var school models.School
	if err := c.ShouldBindJSON(&newSchool); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Bad Parameter"})
		return
	}

	schoolId := c.Param("id")
	object := models.GetDB().First(&school, schoolId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	} else {
		if newSchool.Name != "" {
			object.Update("Name", newSchool.Name)
		}
		c.JSON(http.StatusOK, serializers.SchoolToJSON(school))
	}
}

func DeleteSchool(c *gin.Context) {
	schoolId := c.Param("id")
	var school models.School
	object := models.GetDB().First(&school, schoolId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	} else {
		models.GetDB().Delete(&school, schoolId)
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	}
}
