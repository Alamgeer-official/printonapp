package routes

import (


	"github.com/gin-gonic/gin"

	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middelware"
)

func NewRouter() *gin.Engine {
	// Set Gin mode to release mode
	gin.SetMode(gin.TestMode)

	//init gin engine
	router := gin.Default()

	

	// Setup Middleware
	router.Use(middelware.SetupCORSMiddleware())

	router.GET("", controllers.Test)

	//Login & Signup Routes
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)

	//Admin route
	AdminRoute(router)
	//User Route
	UserRoutes(router)
	//Home page Route
	HomePageSubroute(router)

	return router
}
