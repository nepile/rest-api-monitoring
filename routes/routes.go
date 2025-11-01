package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nepile/api-monitoring/config"
	"github.com/nepile/api-monitoring/controllers"
	"github.com/nepile/api-monitoring/middleware"
)

func Setup(r *gin.Engine, cfg *config.Config) {
	api := r.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.Use(middleware.JWTMiddleware(cfg))
		{
		}
	}
}
