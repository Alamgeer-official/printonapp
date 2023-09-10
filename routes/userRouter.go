package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
)

func UserRoutes(route *gin.Engine) {

	user := route.Group("/users")
	{
		user.GET("/", controllers.GetUser)
	}

}
