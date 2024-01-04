package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middelware"
)

func UserRoutes(route *gin.Engine) {

	user := route.Group("/users")
	{
		user.GET("/", middelware.AuthMiddelware(), controllers.GetUsers)
		user.GET("/user", middelware.AuthMiddelware(), controllers.GetUsers)
	}

}
