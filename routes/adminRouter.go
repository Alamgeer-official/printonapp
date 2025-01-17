package routes

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/controllers"
	"githuh.com/printonapp/middleware"
)

func AdminRoutes(route *gin.Engine) {

	admin := route.Group("/admin", middleware.AuthMiddleware())
	{
		//Products
		productCtr := controllers.NewProductCtr()
		admin.POST("/product", productCtr.AddProducts)

		// Thesis
		thesisCtr := controllers.NewThesisCtr()
		admin.POST("/thesis", thesisCtr.CreateThesis)
		admin.PATCH("/thesis", thesisCtr.UpdateThesisByRole)
		admin.GET("/thesis", thesisCtr.ReadAllThesesByRole)
		admin.GET("/thesis/:id", thesisCtr.GetThesisByID)

		// File upload
		fileCtr := controllers.NewFileController()
		admin.POST("/upload-pdf", fileCtr.UploadPDF)

	}

}
