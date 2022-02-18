package controllers

import (
	"backend/config"
	"backend/forms"
	"backend/models"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
	"time"
)

type Claims struct {
	StudentNumber string `json:"studentNumber"`
	jwt.StandardClaims
}

func Singup(c *gin.Context) {
	var data forms.SignupUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.User{StudentNumber: data.StudentNumber,
		Email:        data.Email,
		Password:     data.Password,
		EntranceYear: data.EntranceYear,
		SchoolId:     data.SchoolId}
	
	if result := models.GetDB().Create(&user); result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return 
	}

	token, err := getToken(user.StudentNumber)
	if err != nil {
		log.Println("err", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "issue in token creation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "user created successfully"})
}

func findUser(studentNumber string) (models.User, error) {
	var user models.User
	if err := models.GetDB().First(&user, "student_number = ?", studentNumber); err.Error != nil {
		return models.User{}, err.Error
	}
	return user, nil
}

func getToken(studentNumber string) (string, error) {
	expirationTime := time.Now().Add(6 * time.Hour)
	claims := &Claims{
		StudentNumber: studentNumber,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	jwtKey := []byte(config.Get("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(c *gin.Context) {
	var data forms.LoginUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := findUser(data.StudentNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !models.CheckPasswordHash(data.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password provided"})
		return
	}
	token, err := getToken(user.StudentNumber)
	if err != nil {
		log.Println("err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "issue in token creation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "message": "user logged in successfully"})
}

func extractToken(c *gin.Context) string {
	bearToken := c.GetHeader("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func extractUserInfo(c *gin.Context) (string, error) {
	tokenString := extractToken(c)
	jwtKey := []byte(config.Get("JWT_SECRET"))
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return claims.StudentNumber, err
}

func JWTAuthenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentNumber, err := extractUserInfo(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		log.Println("user logged in studentNumber:", studentNumber)
	}
}

func GetUserByToken(c *gin.Context) (models.User, error) {
	studentNumber, err := extractUserInfo(c)
	if err != nil {
		return models.User{}, fmt.Errorf("unauthorized")
	}
	user, err := findUser(studentNumber)
	return user, err
}
