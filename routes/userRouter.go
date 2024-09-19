package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middleware"
)

func UserRoutes(route *gin.Engine) {

	mw := middleware.AuthMiddleware()
	user := route.Group("/users")
	{
		user.GET("/", mw, controllers.GetUsers)
		user.GET("/IsEmailExists", controllers.IsEmailExists)

		// Thesis
		thesisCtr := controllers.NewThesisCtr()
		user.GET("/thesis", mw, thesisCtr.ReadAllThesesByRole)
	}

}
