package serializers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

func SchoolToJSON(school models.School) map[string]interface{} {
	return gin.H{"id": school.ID, "name": school.Name}
}

func GetSchoolCourseGroups(school models.School) []map[string]interface{} {
	var groups []models.CourseGroup
	models.GetDB().Preload("Schedule").Find(&groups)

	var courseGroups []map[string]interface{}
	for _, u := range groups {
		var course models.Course
		models.GetDB().First(&course, u.CourseId)
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
