package main

import (
	"backend/config"
	"backend/controllers"
	"backend/models"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config.Init()
	log.Println("Going to initialize Database...")

	DB := models.InitDB()
	controllers.DB = DB
	r := gin.Default()

	conf := cors.DefaultConfig()
	conf.AllowCredentials = true
	conf.AllowAllOrigins = true
	conf.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	conf.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(conf))

	schoolRouter := r.Group("/schools")
	schoolRouter.GET("/", controllers.GetAllSchools)

	authRouter := r.Group("/auth")
	authRouter.POST("/signup", controllers.Singup)
	authRouter.POST("/login", controllers.Login)

	profileRouter := r.Group("/profile")
	profileRouter.Use(controllers.JWTAuthenticator())
	profileRouter.GET("/", controllers.GetProfile)
	profileRouter.PUT("/", controllers.UpdateProfile)

	courseGroupRouter := r.Group("/schools")
	courseGroupRouter.Use(controllers.JWTAuthenticator())
	courseGroupRouter.GET("/course-groups", controllers.GetAllSchoolCourses)

	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
