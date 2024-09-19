package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
)

func HomePageRoutes(route *gin.Engine) {
	homePage := route.Group("/homepage")
	{ // college controller
		collegeCtr := controllers.NewCollegeCtr()
		homePage.GET("/colleges", collegeCtr.GetColleges)

	}
}
