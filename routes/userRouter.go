package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middelware"
)

func UserRoutes(route *gin.Engine) {

	user := route.Group("/users", middelware.AuthMiddelware())
	{
		user.GET("/", controllers.GetUsers)

		// Thesis
		thesisCtr := controllers.NewThesisCtr()
		user.GET("/thesis", thesisCtr.ReadAllThesesByRole)
	}

}
