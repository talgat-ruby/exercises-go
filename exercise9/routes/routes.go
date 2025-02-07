package routes

import (
	"alinurmyrzakhanov/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.AuthMiddleware(), controllers.Logout)

	exerciseGroup := r.Group("/exercises")
	{
		exerciseGroup.POST("", controllers.CreateExercise)
		exerciseGroup.GET("", controllers.GetExercises)
		exerciseGroup.GET("/:id", controllers.GetExerciseByID)
		exerciseGroup.PUT("/:id", controllers.UpdateExercise)
		exerciseGroup.DELETE("/:id", controllers.DeleteExercise)
	}

	authGroup := r.Group("/workouts")
	authGroup.Use(controllers.AuthMiddleware())
	{
		authGroup.POST("", controllers.CreateWorkout)
		authGroup.GET("", controllers.ListWorkouts)
		authGroup.GET("/report", controllers.GetWorkoutReport)
		authGroup.PUT("/:id", controllers.UpdateWorkout)
		authGroup.DELETE("/:id", controllers.DeleteWorkout)
	}
}
