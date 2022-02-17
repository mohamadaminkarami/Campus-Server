package serializers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

func SchoolToJSON(school models.School) map[string]interface{} {
	return gin.H{"id": school.ID, "name": school.Name}
}

func GetSchoolCourseGroups(school models.School, user models.User) []map[string]interface{} {
	var dbCourseGroups []models.CourseGroup
	models.GetDB().Preload("Schedule").Find(&dbCourseGroups)

	var courseGroups []map[string]interface{}
	for _, dbCourseGroup := range dbCourseGroups {
		var course models.Course
		models.GetDB().First(&course, dbCourseGroup.CourseId)
		if course.SchoolId == int(school.ID) {
			courseGroups = append(courseGroups, CourseGroupToJSON(dbCourseGroup, user))
		}
	}
	if courseGroups == nil {
		return []map[string]interface{}{}
	}
	return courseGroups
}

func SchoolCoursesToJSON(school models.School, user models.User) map[string]interface{} {
	return gin.H{
		"schoolId":     school.ID,
		"schoolName":   school.Name,
		"courseGroups": GetSchoolCourseGroups(school, user),
	}
}
