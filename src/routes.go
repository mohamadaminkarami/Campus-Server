package src

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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

type LoginUserData struct {
	StudentNumber string `json:"studentNumber" binding:"required"`
	Password      string `json:"password" binding:"required"`
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
		SchoolId:     data.SchoolId}
	DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"message": "good"})
}

func findUser(studentNumber string) (User, error) {

	var user User
	if err := DB.First(&user, "student_number = ?", studentNumber); err.Error != nil {
		return User{}, err.Error
	}
	return user, nil
}

func Login(c *gin.Context) {
	var data LoginUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := findUser(data.StudentNumber)
	log.Println("uouououo", user.StudentNumber, err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !CheckPasswordHash(data.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password provided"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"message": "should login"})
}
