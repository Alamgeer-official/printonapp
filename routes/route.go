package routes

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	//inir gin engine
	router := gin.Default()
	UserRoutes(router)

	return router
}
