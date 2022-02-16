package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SchoolToJSON(school models.School) map[string]interface{} {
	return gin.H{"ID": school.ID, "name": school.Name}
}

func CreateSchool(c *gin.Context) {
	var School models.School
	if err := c.ShouldBindJSON(&School); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}

	DB.Create(&School)
	c.JSON(http.StatusOK, SchoolToJSON(School))
}

func GetAllSchools(c *gin.Context) {
	var schools []models.School
	DB.Find(&schools)
	var response []map[string]interface{}

	for _, u := range schools {
		response = append(response, SchoolToJSON(u))
	}
	c.JSON(http.StatusOK, response)
}

func UpdateSchool(c *gin.Context) {
	var newSchool models.School
	var school models.School
	if err := c.ShouldBindJSON(&newSchool); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Bad Parameter"})
		return
	}

	schoolId := c.Param("id")
	object := DB.First(&school, schoolId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	} else {
		if newSchool.Name != "" {
			object.Update("Name", newSchool.Name)
		}
		c.JSON(http.StatusOK, SchoolToJSON(school))
	}
}

func DeleteSchool(c *gin.Context) {
	schoolId := c.Param("id")
	var school models.School
	object := DB.First(&school, schoolId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	} else {
		DB.Delete(&school, schoolId)
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	}
}
