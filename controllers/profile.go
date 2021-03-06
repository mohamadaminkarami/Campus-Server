package controllers

import (
	"backend/forms"
	"backend/models"
	"backend/serializers"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func GetProfile(c *gin.Context) {
	user, err := GetUserByToken(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, serializers.UserToJson(user))
}

func getUpdateProfileOmitList(data forms.UpdateUserData) []string {
	omitList := []string{"StudentNumber"}
	for key, element := range structs.Map(data) {
		if element == "" || element == 0 {
			omitList = append(omitList, key)
		}
	}
	return omitList
}

var IsTimestamp validator.Func = func(fl validator.FieldLevel) bool {
	timestamp, ok := fl.Field().Interface().(int)
	if ok {
		timeT, _ := time.Parse("2006-01-02", "2010-01-01")
		if govalidator.IsUnixTime(strconv.Itoa(timestamp)) {
			tm := time.Unix(int64(timestamp), 0)
			return tm.After(timeT)
		}
		return false
	}
	return true
}

var DoesSchoolExist validator.Func = func(fl validator.FieldLevel) bool {
	schoolId, ok := fl.Field().Interface().(int)
	if ok {
		var school models.School
		if result := models.GetDB().First(&school, schoolId); result.Error != nil {
			return false
		}
		return true
	}
	return true
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
	models.GetDB().Model(&user).Omit(getUpdateProfileOmitList(data)...).Updates(structs.Map(data))
	models.GetDB().First(&user, "student_number = ?", stid)
	c.JSON(http.StatusOK, serializers.UserToJson(user))
}
