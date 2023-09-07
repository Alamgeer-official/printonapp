package repository

import (
	"githuh.com/printonapp/models"
)

type UserRepo interface {
	CreateUser(models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetUser() (*[]models.User, error)
}
type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}

}

func (u *userRepo) CreateUser(data models.User) (*models.User, error) {
	result := gormDB.Debug().Create(&data)

	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {

	var user models.User
	res := gormDB.Debug().Where("email", email).Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
func (u *userRepo) GetUser() (*[]models.User, error) {
	var user []models.User
	res := gormDB.Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
