package controllers

import (
	"backend/src"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CourseToJSON(course src.Course) map[string]interface{} {
	return gin.H{
		"ID":            course.ID,
		"name":          course.Name,
		"code":          course.Code,
		"credit":        course.Credit,
		"syllabus":      course.Syllabus,
		"school":        course.SchoolId,
		"prerequisites": course.Prerequisites,
		"corequisites":  course.Corequisites,
	}
}

func CreateCourse(c *gin.Context) {
	var course src.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Error": err})
		return
	}

	var school src.School
	if object := DB.First(&school, course.SchoolId); object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "School not found"})
		return
	}

	DB.Create(&course)
	c.JSON(http.StatusOK, CourseToJSON(course))
}

func GetAllCourses(c *gin.Context) {
	var courses []src.Course
	DB.Preload("Prerequisites").Preload("Corequisites").Find(&courses)
	var response []map[string]interface{}

	for _, u := range courses {
		response = append(response, CourseToJSON(u))
	}
	c.JSON(http.StatusOK, response)
}

func UpdateCourse(c *gin.Context) {
	var newCourse src.Course
	var course src.Course
	if err := c.ShouldBindJSON(&newCourse); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Result": "Bad Parameter"})
		return
	}

	courseId := c.Param("id")
	object := DB.First(&course, courseId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		var prerequisites []src.Course
		for _, pre := range newCourse.Prerequisites {
			var prerequisite src.Course
			if pObject := DB.First(&prerequisite, pre.ID); pObject.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"Error": "Prerequisite not found"})
				return
			} else {
				prerequisites = append(prerequisites, prerequisite)
			}
		}
		DB.Model(&course).Association("Prerequisites").Replace(prerequisites)

		var corequisites []src.Course
		for _, co := range newCourse.Corequisites {
			var corequisite src.Course
			if cObject := DB.First(&corequisite, co.ID); cObject.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"Error": "Corequisite not found"})
				return
			} else {
				corequisites = append(corequisites, corequisite)
			}
		}
		DB.Model(&course).Association("Corequisites").Replace(corequisites)

		if newCourse.Name != "" {
			object.Update("Name", newCourse.Name)
		}
		object.Update("Code", newCourse.Code)
		object.Update("Credit", newCourse.Credit)
		object.Update("Syllabus", newCourse.Syllabus)

		var school src.School
		if sObject := DB.First(&school, course.SchoolId); sObject.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"Error": "School not found"})
			return
		} else {
			object.Update("SchoolId", newCourse.SchoolId)
		}

		c.JSON(http.StatusOK, CourseToJSON(course))
	}
}

func DeleteCourse(c *gin.Context) {
	courseId := c.Param("id")
	var course src.Course
	object := DB.First(&course, courseId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		DB.Delete(&course, courseId)
		c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
	}
}
