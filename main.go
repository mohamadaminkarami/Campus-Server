package main

import (
	"backend/config"
	"backend/controllers"
	"backend/models"
	"flag"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	config.Init()
	log.Println("Starting up...")

	// Initialize Validators
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isTimestamp", controllers.IsTimestamp)
		v.RegisterValidation("doesSchoolExist", controllers.DoesSchoolExist)
		log.Println("Registered validators...")
	}

	// Initialize DB
	models.GetDB()
	log.Println("Initialized Database...")

	// Dummy data
	createDummy := flag.Bool("dummy", false, "insert dummy data. Default is false")
	var users_count int
	flag.IntVar(&users_count, "u", 5, "Specify number of users. Default is 5")
	flag.Parse()
	if *createDummy {
		models.InsertDummyData(users_count)	
		log.Println("Dummy data inserted...")
		// use "go run ." to skip and "go run . -dummy -u Int" to create dummy data
	}
	
	// Routes
	r := gin.Default()

	// CORS
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

	planRouter := r.Group("/plans")
	planRouter.Use(controllers.JWTAuthenticator())
	planRouter.GET("/", controllers.GetAllPlans)
	planRouter.POST("/", controllers.CreatePlan)
	planRouter.DELETE("/:plan_id", controllers.DeletePlan)
	planRouter.GET("/:plan_id", controllers.GetPlan)
	planRouter.POST("/:plan_id/:course_id", controllers.AddCourseToPlan)
	planRouter.DELETE("/:plan_id/:course_id", controllers.DeleteCourseFromPlan)
	planRouter.DELETE("/:plan_id/all", controllers.ClearPlan)

	err := r.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
