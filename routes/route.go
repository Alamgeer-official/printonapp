package routes

import (
	"github.com/gin-gonic/gin"

	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middleware" // Fixed typo in import path
)

// NewRouter initializes a new Gin router with routes and middleware.
func NewRouter() *gin.Engine {
	// Set Gin mode to release mode for production
	gin.SetMode(gin.ReleaseMode)

	// Initialize Gin engine
	router := gin.Default()

	// Setup CORS middleware to handle cross-origin requests
	router.Use(middleware.SetupCORSMiddleware())

	// Define routes
	registerRoutes(router)

	return router
}

// registerRoutes defines all the routes for the application.
func registerRoutes(router *gin.Engine) {
	// Test route to verify server is running
	router.GET("/", controllers.Test)

	// Authentication routes
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)

	// Admin routes
	AdminRoutes(router)

	// User routes
	UserRoutes(router)

	// Home page routes
	HomePageRoutes(router)
	
}