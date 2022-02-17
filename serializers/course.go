package serializers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
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
	course, _ := GetCourseById(courseGroup.CourseId)
	var professor models.Professor
	models.GetDB().First(&professor, courseGroup.ProfessorId)

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
