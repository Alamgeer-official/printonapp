package services

import (
	"time"

	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
)

type UserService interface {
	CreateUser(user models.User) (*models.User, error)
	GetUser() (*[]models.User, error)
}
type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(uRepo repository.UserRepo) UserService {
	return &userService{userRepo: uRepo}

}

func (uSvc *userService) CreateUser(user models.User) (*models.User, error) {
	user.CreatedOn = time.Now()
	data, err := uSvc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uSvc *userService) GetUser() (*[]models.User, error) {
	data, err := uSvc.userRepo.GetUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}
