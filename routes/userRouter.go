package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
)

func UserRoutes(route *gin.Engine) {

	route.POST("/login", controllers.Login)
	route.POST("/signup", controllers.Signup)
	user := route.Group("/users")
	{
		user.GET("/", controllers.GetUser)
	}

}
