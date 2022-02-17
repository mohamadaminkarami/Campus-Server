package controllers

import (
	"backend/forms"
	"backend/models"
	"net/http"

	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func UserToJson(user models.User) map[string]interface{} {
	return gin.H{"studentNumber": user.StudentNumber,
		"email":           user.Email,
		"entranceYear":    user.EntranceYear,
		"takeCoursesTime": user.TakeCoursesTime,
		"SchoolId":        user.SchoolId}
}

func GetProfile(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserToJson(user))
}

func getUpdateProfileOmitList(data forms.UpdateUserData) ([]string) {
	omitList := []string{"StudentNumber"}
	for key, element := range structs.Map(data) {
		if element == "" || element == 0 {
			omitList = append(omitList, key)
		}
    }
	return omitList	
}

func UpdateProfile(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	var data forms.UpdateUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Password != "" {
	hash, err := models.HashPassword(data.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error in hashing"})
		return
	}
		data.Password = hash
	} 
	stid := user.StudentNumber
	DB.Model(&user).Omit(getUpdateProfileOmitList(data)...).Updates(structs.Map(data))
	DB.First(&user, "student_number = ?", stid)
	c.JSON(http.StatusOK, UserToJson(user))
}