package repository

import (
	"errors"

	"githuh.com/printonapp/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserRepo interface {
	CreateUser(models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetUserById(id int64) (*models.User, error)
	GetUsers() (*[]models.User, error)
	IsEmailExists(email string) (bool, error)
}
type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}

}

func (u *userRepo) CreateUser(data models.User) (*models.User, error) {
	result := gormDB.Create(&data)

	if result.Error != nil {
		return nil, result.Error
	}
	return &data, nil
}

func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {

	var user models.User
	res := gormDB.Preload(clause.Associations).Where("email", email).Where("active=true").Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
func (u *userRepo) GetUserById(id int64) (*models.User, error) {

	var user models.User
	res := gormDB.Where("id", id).Where("active=true").Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}

	return &user, nil
}
func (u *userRepo) GetUsers() (*[]models.User, error) {
	var user []models.User
	res := gormDB.Omit("password").Where("active = true").Where("role", "USER").Find(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

func (u *userRepo) IsEmailExists(email string) (bool, error) {
	var user models.User
	res := gormDB.Where("email = ?", email).First(&user)
	if res.Error != nil && errors.Is(res.Error, gorm.ErrRecordNotFound) {
		// No record found, email does not exist
		return false, nil
	} else if res.Error != nil {
		// Some other error occurred
		return false, res.Error
	}
	// Record found, email exists
	return true, nil
}
