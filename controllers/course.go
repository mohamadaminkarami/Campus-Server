package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CourseToJSON(course models.Course) map[string]interface{} {
	return gin.H{
		"id":            course.ID,
		"name":          course.Name,
		"code":          course.Code,
		"credit":        course.Credit,
		"syllabus":      course.Syllabus,
		"school":        course.SchoolId,
		"prerequisites": RequisitesToJSON(course.Prerequisites),
		"corequisites":  RequisitesToJSON(course.Corequisites),
	}
}

func ProfessorToJSON(professor models.Professor) map[string]interface{} {
	return gin.H{
		"id":   professor.ID,
		"name": professor.Name,
	}
}

func CourseGroupToJSON(courseGroup models.CourseGroup) map[string]interface{} {
	var course models.Course
	DB.Preload("Prerequisites").Preload("Corequisites").First(&course, courseGroup.CourseId)
	var professor models.Professor
	DB.First(&professor, courseGroup.ProfessorId)

	return gin.H{
		"id":          courseGroup.ID,
		"professor":   ProfessorToJSON(professor),
		"course":      CourseToJSON(course),
		"groupNumber": courseGroup.GroupNumber,
		"capacity":    courseGroup.Capacity,
		"examDate":    courseGroup.ExamDate.Unix(),
		"detail":      courseGroup.Detail,
		"schedule":    SchedulesToJSON(courseGroup.Schedule),
	}
}

func SchedulesToJSON(schedules []models.Schedule) []map[string]interface{} {
	var response []map[string]interface{}
	for _, schedule := range schedules {
		response = append(response, gin.H{
			"startTime": schedule.Start,
			"endTime":   schedule.End,
			"day":       schedule.Day,
		})
	}
	if response == nil {
		return []map[string]interface{}{}
	}
	return response
}

func RequisitesToJSON(courses []models.Course) []map[string]interface{} {
	var response []map[string]interface{}
	for _, course := range courses {
		response = append(response, gin.H{
			"id":     course.ID,
			"name":   course.Name,
			"code":   course.Code,
			"credit": course.Credit,
		})
	}
	if response == nil {
		return []map[string]interface{}{}
	}
	return response
}

func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"Error": err})
		return
	}

	var school models.School
	if object := DB.First(&school, course.SchoolId); object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "School not found"})
		return
	}

	DB.Create(&course)
	c.JSON(http.StatusOK, CourseToJSON(course))
}

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	DB.Preload("Prerequisites").Preload("Corequisites").Find(&courses)
	var response []map[string]interface{}

	for _, u := range courses {
		response = append(response, CourseToJSON(u))
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
	object := DB.First(&course, courseId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		var prerequisites []models.Course
		for _, pre := range newCourse.Prerequisites {
			var prerequisite models.Course
			if pObject := DB.First(&prerequisite, pre.ID); pObject.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{"Error": "Prerequisite not found"})
				return
			} else {
				prerequisites = append(prerequisites, prerequisite)
			}
		}
		DB.Model(&course).Association("Prerequisites").Replace(prerequisites)

		var corequisites []models.Course
		for _, co := range newCourse.Corequisites {
			var corequisite models.Course
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

		var school models.School
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
	var course models.Course
	object := DB.First(&course, courseId)

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Item not found"})
	} else {
		DB.Delete(&course, courseId)
		c.JSON(http.StatusOK, gin.H{"Success": "Item deleted"})
	}
}
