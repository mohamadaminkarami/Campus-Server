package main

import (
	"backend/controllers"
	. "backend/src"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Going to initialize Database...")

	var db gorm.DB
	db = InitDB()
	controllers.DB = db
	result := db.Create(&User{StudentNumber: "98101244", Password: "password", Email: "masihbr@gmail.com", EntranceYear: 1398})
	log.Println(result)
	var user User
	db.First(&user, "student_number = ?", "98101244")
	log.Println("DB find: ", user.StudentNumber, user.Password)

	r := gin.Default()
	r.GET("/ping", Pong)

	schoolRouter := r.Group("/schools")
	schoolRouter.POST("/", controllers.CreateSchool)
	schoolRouter.PUT("/:id", controllers.UpdateSchool)
	schoolRouter.DELETE("/:id", controllers.DeleteSchool)
	schoolRouter.GET("/", controllers.GetAllSchools)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
