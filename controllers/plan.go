package controllers

import (
	"backend/models"
	"backend/serializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePlan(c *gin.Context) {
	var plan models.Plan
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	plan.UserId = int(user.ID)
	models.GetDB().Create(&plan)
	c.JSON(http.StatusOK, serializers.PlanToJSON(plan))
}

func GetAllPlans(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	var plans []models.Plan
	models.GetDB().Where(&models.Plan{UserId: int(user.ID)}).Find(&plans)
	var response []map[string]interface{}

	for _, u := range plans {
		response = append(response, serializers.PlanToJSON(u))
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
	result := models.GetDB().Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
	} else {
		models.GetDB().Delete(&plan, planId)
		c.JSON(http.StatusOK, gin.H{"message": "plan deleted"})
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
	result := models.GetDB().Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
	} else {
		c.JSON(http.StatusOK, serializers.PlanToJSON(plan))
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
	result := models.GetDB().Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
		return
	}

	courseGroupId := c.Param("course_id")
	var courseGroup models.CourseGroup
	result = models.GetDB().First(&courseGroup, courseGroupId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "courseGroup not found"})
		return
	}

	err = models.GetDB().Model(&plan).Association("Courses").Append([]models.CourseGroup{courseGroup})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializers.PlanToJSON(plan))
}

func DeleteCourseFromPlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := models.GetDB().Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plan not found"})
		return
	}

	courseGroupId := c.Param("course_id")
	var courseGroup models.CourseGroup
	result = models.GetDB().First(&courseGroup, courseGroupId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "CourseGroup not found"})
		return
	}

	err = models.GetDB().Model(&plan).Association("Courses").Delete([]models.CourseGroup{courseGroup})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializers.PlanToJSON(plan))
}

func ClearPlan(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	planId := c.Param("plan_id")
	var plan models.Plan
	result := models.GetDB().Where(&models.Plan{UserId: int(user.ID)}).First(&plan, planId)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "plan not found"})
		return
	}

	err = models.GetDB().Model(&plan).Association("Courses").Clear()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializers.PlanToJSON(plan))
}
