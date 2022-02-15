package main

import (
	"backend/controllers"
	. "backend/src"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println("Going to initialize Database...")

	DB = InitDB()
	controllers.DB = DB
	DB.Create(&School{Name: "CE"})
	result := DB.Create(&User{StudentNumber: "98101244", Password: "password", Email: "masihbr@gmail.com", EntranceYear: 1398, SchoolId: 1})
	log.Println(result)
	var user User
	DB.First(&user, "student_number = ?", "98101244")
	log.Println("DB find: ", user.StudentNumber, user.Password)

	r := gin.Default()
	r.GET("/ping", Pong)
	
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	schoolRouter := r.Group("/schools")
	schoolRouter.POST("/", controllers.CreateSchool)
	schoolRouter.PUT("/:id", controllers.UpdateSchool)
	schoolRouter.DELETE("/:id", controllers.DeleteSchool)
	schoolRouter.GET("/", controllers.GetAllSchools)

	authRouter := r.Group("/auth")
	authRouter.POST("/signup", Singup)
	authRouter.POST("/login", Login)

	profileRouter := r.Group("/profile")
	profileRouter.Use(JWTAuthenticator())
	profileRouter.GET("/", GetProfile)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
