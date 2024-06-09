package routes

import (
	"jwt-auth-service/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
}
