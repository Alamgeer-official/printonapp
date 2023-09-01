package repository

import "githuh.com/printonapp/models"

type UserRepo interface {
	CreateUser(data models.User) (*models.User, error)
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

func (u *userRepo) GetUser() (*[]models.User, error) {
	var user []models.User
	res := gormDB.Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}
