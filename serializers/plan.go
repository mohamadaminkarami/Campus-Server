package serializers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

func PlanToJSON(plan models.Plan) map[string]interface{} {
	var courses []models.CourseGroup
	err := models.GetDB().Preload("Schedule").Model(&plan).Association("Courses").Find(&courses)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return gin.H{
		"id":           plan.ID,
		"userId":       plan.UserId,
		"totalCredits": GetTotalCredits(courses),
		"courseGroups": GetPlanCourseGroups(courses),
	}
}

func GetPlanCourseGroups(courses []models.CourseGroup) []int {
	var courseGroups []int
	for _, u := range courses {
		courseGroups = append(courseGroups, int(u.ID))
	}
	if courseGroups == nil {
		return []int{}
	}
	return courseGroups
}

func GetTotalCredits(courses []models.CourseGroup) int {
	sum := 0
	for _, courseGroup := range courses {
		course, _ := GetCourseById(courseGroup.CourseId)
		sum += course.Credit
	}
	return sum
}

func GetCourseById(courseId int) (models.Course, error) {
	var course models.Course
	result := models.GetDB().Preload("Prerequisites").Preload("Corequisites").First(&course, courseId)

	if result.Error != nil {
		return models.Course{}, result.Error
	}
	return course, nil
}

func GetProfessorById(professorId int) (models.Professor, error) {
	var professor models.Professor
	result := models.GetDB().First(&professor, professorId)

	if result.Error != nil {
		return models.Professor{}, result.Error
	}
	return professor, nil
}
