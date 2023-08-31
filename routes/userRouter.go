package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
)

func UserRoutes(route *gin.Engine) {
	user := route.Group("users")
	{
		user.GET("/user2", controllers.UserTest2)
		user.GET("/user", controllers.UserTest)
	}

}
