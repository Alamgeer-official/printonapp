package repository

import "githuh.com/printonapp/models"

type ProductRepo interface {
	CreateProduct(product models.Product) (string, error)
}

type productRepo struct{}

func NewProductRepo() ProductRepo {
	return &productRepo{}
}

func (p *productRepo) CreateProduct(product models.Product) (string, error) {
	res := gormDB.Create(&product)
	if res.Error != nil {
		return "", res.Error
	}
	return "Created", nil
}
