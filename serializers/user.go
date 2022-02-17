package serializers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
)

func UserToJson(user models.User) map[string]interface{} {
	return gin.H{"studentNumber": user.StudentNumber,
		"email":           user.Email,
		"entranceYear":    user.EntranceYear,
		"takeCoursesTime": user.TakeCoursesTime,
		"SchoolId":        user.SchoolId}
}
