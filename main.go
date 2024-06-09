package main

import (
	"jwt-auth-service/config"
	"jwt-auth-service/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()

	// Authentication routes
	authRoutes := r.Group("/auth")
	routes.AuthRoutes(authRoutes)

	// User routes
	userRoutes := r.Group("/users")
	routes.UserRoutes(userRoutes)

	r.Run(":8080")
}
