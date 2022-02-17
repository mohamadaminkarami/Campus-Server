package controllers

import (
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PlanToJSON(plan models.Plan) map[string]interface{} {
	var courses []models.CourseGroup
	err := DB.Preload("Schedule").Model(&plan).Association("Courses").Find(&courses)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return gin.H{
		"id":      plan.ID,
		"userId":  plan.UserId,
		"totalCredits": GetTotalCredits(courses),
		"courses": GetPlanCourseGroups(courses),
	}
}

func GetPlanCourseGroups(courses []models.CourseGroup) []map[string]interface{} {
	var courseGroups []map[string]interface{}
	for _, u := range courses {
		courseGroups = append(courseGroups, CourseGroupToJSON(u))
	}
	if courseGroups == nil {
		return []map[string]interface{}{}
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

func CreatePlan(c *gin.Context) {
	var plan models.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	plan.UserId = int(user.ID)
	DB.Create(&plan)
	c.JSON(http.StatusOK, PlanToJSON(plan))
}

func GetAllPlans(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	
	var plans []models.Plan
	DB.Where(&models.Plan{UserId: int(user.ID)}).Find(&plans)
	var response []map[string]interface{}

	for _, u := range plans {
		response = append(response, PlanToJSON(u))
	}
	c.JSON(http.StatusOK, gin.H{"result": response})
}

func DeletePlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := DB.Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	} else {
		DB.Delete(&plan, planId)
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	}
}

func GetPlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := DB.Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
	} else {
		c.JSON(http.StatusOK, PlanToJSON(plan))
	}
}

func AddCourseToPlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := DB.Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId) // TODO: Filter By User
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	courseGroupId := c.Param("course_id")
	var courseGroup models.CourseGroup
	result = DB.First(&courseGroup, courseGroupId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	err = DB.Model(&plan).Association("Courses").Append([]models.CourseGroup{courseGroup})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, PlanToJSON(plan))
}

func DeleteCourseFromPlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := DB.Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId) // TODO: Filter By User
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	courseGroupId := c.Param("course_id")
	var courseGroup models.CourseGroup
	result = DB.First(&courseGroup, courseGroupId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	err = DB.Model(&plan).Association("Courses").Delete([]models.CourseGroup{courseGroup})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, PlanToJSON(plan))
}

func ClearPlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := DB.Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId) // TODO: Filter By User
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	err = DB.Model(&plan).Association("Courses").Clear()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, PlanToJSON(plan))
}