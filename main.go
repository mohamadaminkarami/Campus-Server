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
	authRouter.POST("/signup", controllers.Singup)
	authRouter.POST("/login", controllers.Login)

	profileRouter := r.Group("/profile")
	profileRouter.Use(controllers.JWTAuthenticator())
	profileRouter.GET("/", controllers.GetProfile)
	profileRouter.PUT("/", controllers.UpdateProfile)

	courseRouter := r.Group("/courses")
	courseRouter.POST("/", controllers.CreateCourse)
	courseRouter.PUT("/:id", controllers.UpdateCourse)
	courseRouter.DELETE("/:id", controllers.DeleteCourse)
	courseRouter.GET("/", controllers.GetAllCourses)

	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
