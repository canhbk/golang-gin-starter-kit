package routes

import (
	"github.com/canhbk/golang-gin-starter-kit/controllers"
	v1 "github.com/canhbk/golang-gin-starter-kit/controllers/v1"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	// API Version 1 Routes
	v1Routes := r.Group("/api/v1")
	initializeV1Routes(v1Routes)

	// Health check route (unversioned)
	healthController := controllers.NewHealthController()
	r.GET("/health", healthController.HealthCheck)
}

func initializeV1Routes(rg *gin.RouterGroup) {
	// Initialize V1 controllers
	userController := v1.NewUserController()

	// User routes
	users := rg.Group("/users")
	{
		users.POST("", userController.Create)
		users.GET("", userController.List)
		users.GET("/:id", userController.Get)
		users.PUT("/:id", userController.Update)
		users.DELETE("/:id", userController.Delete)
	}

	// Add other v1 route groups here
}

// Prepare for future versions
func initializeV2Routes(rg *gin.RouterGroup) {
	// V2 routes will go here
}
