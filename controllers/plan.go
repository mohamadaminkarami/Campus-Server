package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PlanToJSON(plan models.Plan) map[string]interface{} {
	return gin.H{
		"id":      plan.ID,
		"userId":  plan.UserId,
		"courses": plan.Courses,
	}
}

func CreatePlan(c *gin.Context) {
	var plan models.Plan
	if err := c.ShouldBindJSON(&plan); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	// TODO: Set userId
	DB.Create(&plan)
	c.JSON(http.StatusOK, PlanToJSON(plan))
}

func GetAllPlans(c *gin.Context) {
	var plans []models.Plan
	DB.Find(&plans) // TODO: Filter by User
	var response []map[string]interface{}

	for _, u := range plans {
		response = append(response, PlanToJSON(u))
	}
	c.JSON(http.StatusOK, gin.H{"result": response})
}

func DeletePlan(c *gin.Context) {
	planId := c.Param("id")
	var plan models.Plan
	object := DB.First(&plan, planId) // TODO: Filter By User

	if object.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
	} else {
		DB.Delete(&plan, planId)
		c.JSON(http.StatusOK, gin.H{"message": "Item deleted"})
	}
}
