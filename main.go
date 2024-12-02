package main

import (
	"college-be/controllers"
	"college-be/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "POST", "GET", "OPTIONS"},
		AllowHeaders:  []string{"Content-Type", "X-Requested-With", "Authorization"},
		ExposeHeaders: []string{"Content-Length"},
	}))
	godotenv.Load()

	database.Connect()
	database.Migrate()
	api := r.Group("/api")
	api.POST("/auth/register", controllers.RegisterUser)
	api.POST("/auth/login", controllers.GenerateToken)
	api.POST("/mahasiswa/:nim/matakuliah/:courseId", controllers.AddCourse)
	api.DELETE("/mahasiswa/:nim/matakuliah/:courseId", controllers.RemoveCourse)

	api.GET("/majors", controllers.GetAllMajors)
	api.GET("/mahasiswa", controllers.GetAllUsers)
	api.GET("/mahasiswa/:nim", controllers.GetUserByNIM)
	api.PUT("/mahasiswa/profile", controllers.UpdateUserData)
	api.GET("/mahasiswa/profile", controllers.GetUserProfile)
	api.GET("/matakuliah", controllers.GetAllCourses)

	r.Run()
}
