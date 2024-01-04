package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middelware"
)

func NewRouter() *gin.Engine {
	//init gin engine
	router := gin.Default()
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
