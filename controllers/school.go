package controllers

import (
	"backend/src"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

var DB gorm.DB

func SchoolToJSON(school src.School) map[string]interface{} {
	return gin.H{"ID": school.ID, "name": school.Name}
}

func CreateSchool(c *gin.Context) {
	var School src.School
	if err := c.ShouldBindJSON(&School); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Error": err})
		return
	}

	DB.Create(&School)
	c.JSON(http.StatusOK, SchoolToJSON(School))
}

func GetAllSchools(c *gin.Context) {
	var schools []src.School
	DB.Find(&schools)
	var response []map[string]interface{}

	for _, u := range schools {
		response = append(response, SchoolToJSON(u))
	}
	c.JSON(http.StatusOK, response)
}

func UpdateSchool(c *gin.Context) {
	var newSchool src.School
	var school src.School
	if err := c.ShouldBindJSON(&newSchool); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
		return
	}

	schoolId := c.Param("id")
	object := DB.First(&school, schoolId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		if newSchool.Name != "" {
			object.Update("Name", newSchool.Name)
		}
		c.JSON(http.StatusOK, SchoolToJSON(school))
	}
}

func DeleteSchool(c *gin.Context) {
	schoolId := c.Param("id")
	var school src.School
	object := DB.First(&school, schoolId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		DB.Delete(&school, schoolId)
		c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
	}
}
