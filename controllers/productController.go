package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/services"
	"githuh.com/printonapp/utils"
)

type ProductCtr interface {
	AddProducts(ctx *gin.Context)
}
type productCtr struct{}

func NewProductCtr() ProductCtr {
	return &productCtr{}
}

var productSrv = services.NewProuductSrv(repository.NewProductRepo())

func (pCtr *productCtr) AddProducts(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBind(&product); err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}

	// call create product service func
	data, err := productSrv.CreateProduct(ctx, product)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)

}
