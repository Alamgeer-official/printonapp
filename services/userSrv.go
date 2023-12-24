package services

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(models.User) (*models.User, error)
	Login(map[string]string) (*models.User, error)
	GetUserByEmail(string) (*models.User, error)
	GetUsers(ctx *gin.Context) (*[]models.User, error)
}
type userService struct {
	userRepo repository.UserRepo
}

func NewUserService(uRepo repository.UserRepo) UserService {
	return &userService{userRepo: uRepo}

}

func (uSvc *userService) CreateUser(user models.User) (*models.User, error) {
	// validation
	if user.FirstName == "" || user.Email == "" || user.Phone == "" || user.Password == "" || user.Role == "" {
		return nil, errors.New("mandatory feild is empty")
	}
	// Get user detail
	userData, err := uSvc.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	}
	//validate already exist
	if userData.Email == user.Email {
		return nil, fmt.Errorf("user already exist : %s", user.Email)
	}
	//encrypt password
	pwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// assign values
	user.Password = string(pwd)
	user.CreatedOn = time.Now()
	user.Email = strings.ToLower(user.Email)
	user.Active = true
	user.Role = "USER"
	data, err := uSvc.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uSvc *userService) Login(credential map[string]string) (*models.User, error) {
	//validations
	userData, err := uSvc.GetUserByEmail(credential["email"])
	if err != nil {
		return nil, err
	}
	if userData.Email == "" {
		return nil, fmt.Errorf("user not found : %s", userData.Email)
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(credential["password"]))
	if err != nil {
		return nil, err // Password mismatch or invalid hash
	}

	accesToken, err := utils.CreateJWToken(userData)
	if err != nil {
		return nil, err
	}

	userData.AccessToken = accesToken
	//omit password
	userData.Password = ""
	return userData, nil
}

func (uSvc *userService) GetUserByEmail(email string) (*models.User, error) {
	email = strings.ToLower(email)
	data, err := uSvc.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	return data, nil
}
func (uSvc *userService) GetUsers(ctx *gin.Context) (*[]models.User, error) {
	user := utils.GetUserDataFromContext(ctx)
	if !user.IsAdmin() {
		return nil, errors.New("only admin allowed")
	}

	data, err := uSvc.userRepo.GetUsers()
	if err != nil {
		return nil, err
	}
	return data, nil
}
