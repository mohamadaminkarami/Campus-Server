package src

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

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

type Claims struct {
	StudentNumber string `json:"studentNumber"`
	jwt.StandardClaims
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
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Login(c *gin.Context) {
	var data LoginUserData
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := findUser(data.StudentNumber)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if !CheckPasswordHash(data.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong password provided"})
		return
	}
	token, err := getToken(user.StudentNumber)
	if err != nil {
		log.Println("err", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "issue in token creation"})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"access": token})
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
	log.Println("token string:", tokenString)
	var claims Claims
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	log.Println("student:", claims.StudentNumber)
	return claims.StudentNumber, err
}

func JWTAuthenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		studentNumber, err := extractUserInfo(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		log.Println("studentNumber:", studentNumber)
	}
}

func GetProfile(c *gin.Context) {
	studentNumber, err := extractUserInfo(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
		return
	}
	var user User
	if err := DB.First(&user, "student_number = ?", studentNumber); err.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user specified by token not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"studentNumber": user.StudentNumber, "entranceYear": user.EntranceYear})
}
