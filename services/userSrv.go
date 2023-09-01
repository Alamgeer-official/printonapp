package services

import (
	"time"

	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
)

func CreateUser(user models.User) (*models.User, error) {
	user.CreatedOn = time.Now()
	data, err := repository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUser() (*[]models.User, error) {
	data, err := repository.GetUser()
	if err != nil {
		return nil, err
	}
	return data, nil
}
