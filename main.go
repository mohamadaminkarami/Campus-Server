package main

import (
	src "backend/src"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"fmt"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
	  
	fmt.Println("Going to initialize Database...")

	var db gorm.DB
	db = src.InitDB()
	// result := db.Create(&User{StudentNumber: "98101244", Password: "password", Email: "masihbr@gmail.com", EntranceYear: 1398})
	// log.Println(result)
	var user src.User
	db.First(&user, "student_number = ?", "98101244")
	log.Println("DB find: ", user.StudentNumber, user.Password)

	r := gin.Default()
	r.GET("/ping", src.Pong)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
