package repository

import "githuh.com/printonapp/models"

var (
	user = []models.User{}
)

func CreateUser(data *models.User) (*models.User, error) {
	user = append(user, *data)
	return data, nil
}

func GetUser() (*[]models.User, error) {
	return &user, nil
}
