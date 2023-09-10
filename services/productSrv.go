package services

import (
	"errors"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/utils"
)

type ProductSrv interface {
	CreateProduct( ctx *gin.Context,product models.Product,) (string, error)
}
type productSrv struct {
	productRepo repository.ProductRepo
}

func NewProuductSrv(pRepo repository.ProductRepo) ProductSrv {
	return &productSrv{productRepo: pRepo}
}

func (pSrv *productSrv) CreateProduct( ctx *gin.Context,product models.Product,) (string, error) {
	user := utils.GetUserDataFromContext(ctx)
	if !user.IsAdmin(){
		return"",errors.New("unauthorized access")
	}
	if product.Name==""{
		return"",errors.New("invalid product name")
	}
	product.Active=true
	product.UserID=user.ID
	res, err := pSrv.productRepo.CreateProduct(product)
	if err != nil {
		return "", err
	}
	return res, nil

}
