package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePageSubroute(route *gin.Engine) {
	homePage := route.Group("/homepage")
	{
		homePage.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "im in homepage"})

		})
		homePage.GET("/test", func(ctx *gin.Context) {
			ctx.IndentedJSON(http.StatusAccepted, "i m in test home page")
		})
	}
}
