package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
)

func UserRoutes(route *gin.Engine) {

	route.GET("/login", controllers.Login)
	route.POST("/signup", controllers.Signup)
	// user := route.Group("/user")
	// {
	// 	user.GET("/signup", controllers.Signup)
	// 	user.GET("/login", controllers.Login)
	// }

}
