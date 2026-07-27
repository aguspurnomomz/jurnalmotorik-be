package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/aguspurnomomz/journalmotorik-app-be/controllers"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		// Auth Routes
		auth := api.Group("/auth")
		{
			auth.POST("/register", controllers.Register)
			auth.POST("/login", controllers.Login)
		}
	}
}