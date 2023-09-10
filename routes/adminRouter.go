package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middelware"
)

func AdminRoute(route *gin.Engine) {

	admin := route.Group("/admin", middelware.AuthMiddelware())
	{

		//Products
		productCtr := controllers.NewProductCtr()
		admin.POST("/product", productCtr.AddProducts)

	}
}
