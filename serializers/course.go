package serializers

import (
	"backend/models"
	"math/rand"

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

func indexOf(element int, data []int) (int) {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 
 }

func CalcTakeChance( courseGroup models.CourseGroup, user models.User) (int) {
	/*
	capacity := courseGroup.Capacity
	var users []models.User
	var selectedIds []int
	models.GetDB().Where("take_courses_time > 0").Order("take_courses_time").Find(&users)
	for _, user := range users {
	  var plan models.Plan
	  models.GetDB().Where("user_id = ?", user.ID).First(&plan)
	  var courseGroups []models.CourseGroup
	  models.GetDB().Model(&plan).Association("Courses").Find(&courseGroups)
	  for _, dbCourseGroup := range courseGroups {
		if dbCourseGroup.ID == courseGroup.ID {
		  selectedIds = append(selectedIds, int(user.ID))
		  break
		}
	  }
	}
	index := indexOf(int(user.ID), selectedIds)
	if (index - capacity) <= 0 {
	  rand.Seed(time.Now().UnixNano())
	  return rand.Intn(50) + 50
	}
	log.Println(courseGroup.CourseId, index, capacity, len(selectedIds))  
	return int((1.05 - float64(index - capacity) / float64(len(selectedIds) - capacity) ) * 100)
	*/
	return rand.Intn(50)
}

func CourseGroupToJSON(courseGroup models.CourseGroup, user models.User) map[string]interface{} {
	course, _ := GetCourseById(courseGroup.CourseId)
	professor, _ := GetProfessorById(courseGroup.ProfessorId)

	return gin.H{
		"id":          courseGroup.ID,
		"professor":   ProfessorToJSON(professor),
		"course":      CourseToJSON(course),
		"groupNumber": courseGroup.GroupNumber,
		"capacity":    courseGroup.Capacity,
		"examDate":    courseGroup.ExamDate.Unix(),
		"detail":      courseGroup.Detail,
		"schedule":    SchedulesToJSON(courseGroup.Schedule),
		"takeChance":  CalcTakeChance(courseGroup, user),
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
