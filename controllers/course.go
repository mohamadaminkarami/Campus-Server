package controllers

import (
	"backend/models"
	"backend/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Error": err})
		return
	}

	var school models.School
	if object := models.GetDB().First(&school, course.SchoolId); object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "School not found"})
		return
	}

	models.GetDB().Create(&course)
	c.JSON(http.StatusOK, serializers.CourseToJSON(course))
}

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	models.GetDB().Preload("Prerequisites").Preload("Corequisites").Find(&courses)
	var response []map[string]interface{}

	for _, u := range courses {
		response = append(response, serializers.CourseToJSON(u))
	}
	c.JSON(http.StatusOK, response)
}

func UpdateCourse(c *gin.Context) {
	var newCourse models.Course
	var course models.Course
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
		return
	}

	courseId := c.Param("id")
	object := models.GetDB().First(&course, courseId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		var prerequisites []models.Course
		for _, pre := range newCourse.Prerequisites {
			var prerequisite models.Course
			if pObject := models.GetDB().First(&prerequisite, pre.ID); pObject.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"Error": "Prerequisite not found"})
				return
			} else {
				prerequisites = append(prerequisites, prerequisite)
			}
		}
		models.GetDB().Model(&course).Association("Prerequisites").Replace(prerequisites)

		var corequisites []models.Course
		for _, co := range newCourse.Corequisites {
			var corequisite models.Course
			if cObject := models.GetDB().First(&corequisite, co.ID); cObject.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"Error": "Corequisite not found"})
				return
			} else {
				corequisites = append(corequisites, corequisite)
			}
		}
		models.GetDB().Model(&course).Association("Corequisites").Replace(corequisites)

		if newCourse.Name != "" {
			object.Update("Name", newCourse.Name)
		}
		object.Update("Code", newCourse.Code)
		object.Update("Credit", newCourse.Credit)
		object.Update("Syllabus", newCourse.Syllabus)

		var school models.School
		if sObject := models.GetDB().First(&school, course.SchoolId); sObject.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "School not found"})
			return
		} else {
			object.Update("SchoolId", newCourse.SchoolId)
		}

		c.JSON(http.StatusOK, serializers.CourseToJSON(course))
	}
}

func DeleteCourse(c *gin.Context) {
	courseId := c.Param("id")
	var course models.Course
	object := models.GetDB().First(&course, courseId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		models.GetDB().Delete(&course, courseId)
		c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
	}
}
