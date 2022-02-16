package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SchoolToJSON(school models.School) map[string]interface{} {
	return gin.H{"id": school.ID, "name": school.Name}
}

func GetSchoolCourseGroups(school models.School) []map[string]interface{} {
	var groups []models.CourseGroup
	DB.Preload("Schedule").Find(&groups)

	var courseGroups []map[string]interface{}
	for _, u := range groups {
		var course models.Course
		DB.First(&course, u.CourseId)
		if course.SchoolId == int(school.ID) {
			courseGroups = append(courseGroups, CourseGroupToJSON(u))
		}
	}
	if courseGroups == nil {
		return []map[string]interface{}{}
	}
	return courseGroups
}

func SchoolCoursesToJSON(school models.School) map[string]interface{} {
	return gin.H{
		"schoolId":     school.ID,
		"schoolName":   school.Name,
		"courseGroups": GetSchoolCourseGroups(school),
	}
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

func GetAllSchoolCourses(c *gin.Context) {
	var schools []models.School
	DB.Find(&schools)
	var response []map[string]interface{}

	for _, u := range schools {
		response = append(response, SchoolCoursesToJSON(u))
	}
	c.JSON(http.StatusOK, gin.H{"result": response})
}

func GetAllSchools(c *gin.Context) {
	var schools []models.School
	DB.Find(&schools)
	var response []map[string]interface{}

	for _, u := range schools {
		response = append(response, SchoolToJSON(u))
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
