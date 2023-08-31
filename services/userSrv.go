package services

import (
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
)

func CreateUser(user *models.User) (*models.User, error) {
	data, _ := repository.CreateUser(user)
	return data, nil
}

func GetUser() (*[]models.User, error) {
	data, _ := repository.GetUser()
	return data, nil
}
