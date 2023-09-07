package services

import (
	"strings"
	"time"

	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
)

type UserService interface {
	CreateUser(models.User) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
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
	user.Email = strings.ToLower(user.Email)
	user.Active=true
	data, err := uSvc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (uSvc *userService) GetUserByEmail(email string) (*models.User, error) {
	email=strings.ToLower(email)
	data, err := uSvc.userRepo.GetUserByEmail(email)
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
