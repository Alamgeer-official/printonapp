package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middelware"
)

func UserRoutes(route *gin.Engine) {

	route.POST("/login", controllers.Login)
	route.POST("/signup", controllers.Signup)
	user := route.Group("/users", middelware.AuthMiddelware())
	{
		user.GET("/", controllers.GetUser)
	}

}
