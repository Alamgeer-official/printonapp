package routes

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	//inir gin engine
	router := gin.Default()

	//User Route
	UserRoutes(router)

	//Home page Route
	HomePageSubroute(router)
	return router
}
