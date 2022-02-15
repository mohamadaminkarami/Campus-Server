package src

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	log.Println("ping requested...")
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

type SignupUserData struct {
	// binding:"required" ensures that the field is provided
	StudentNumber string `json:"studentNumber" binding:"required"`
	Email         string `json:"email" binding:"required"`
	Password      string `json:"password" binding:"required"`
	EntranceYear  int    `json:"entranceYear" binding:"required"`
	SchoolId      int    `json:"schoolId" binding:"required"`
}

func Singup(c *gin.Context) {
	var data SignupUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println("valid input")
	user := User{StudentNumber: data.StudentNumber,
		Email:        data.Email,
		Password:     data.Password,
		EntranceYear: data.EntranceYear,
		SchoolId:      data.SchoolId}
	DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "good"})
}
