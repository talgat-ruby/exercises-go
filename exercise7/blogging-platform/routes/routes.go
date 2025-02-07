// routes/routes.go
package routes

import (
	"alinurmyrzakhanov/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	postRoutes := r.Group("/posts")
	{
		postRoutes.POST("", controllers.CreatePost)
		postRoutes.GET("", controllers.GetAllPosts)
		postRoutes.GET("/:id", controllers.GetPost)
		postRoutes.PUT("/:id", controllers.UpdatePost)
		postRoutes.DELETE("/:id", controllers.DeletePost)
	}
}
