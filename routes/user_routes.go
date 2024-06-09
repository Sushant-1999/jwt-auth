package routes

import (
	"jwt-auth-service/controllers"
	"jwt-auth-service/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup) {
	r.Use(middleware.AuthMiddleware())
	r.GET("/", controllers.GetUsers)
	r.GET("/:id", controllers.GetUserByID)
}
