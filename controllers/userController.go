package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/services"
	"githuh.com/printonapp/utils"
)

var userService = services.NewUserService(repository.NewUserRepo())

func Signup(ctx *gin.Context) {

	// bind signup data in user model
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	// validation
	if user.FirstName == "" || user.Email == "" || user.Phone == "" || user.Password == "" || user.Role == "" {
		utils.ReturnError(ctx, errors.New("mandatory feild is empty"), http.StatusBadRequest)
		return
	}
	userData, err := userService.GetUserByEmail(user.Email)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	if userData.Email == user.Email {
		utils.ReturnError(ctx, errors.New("account already created"), http.StatusConflict)
		return
	}

	//create user
	data, err := userService.CreateUser(user)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusConflict)
		return
	}

	ctx.JSON(http.StatusCreated, data)
}
func Login(ctx *gin.Context) {
	var user map[string]string
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	userData, err := userService.GetUserByEmail(user["email"])
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}
	if userData.Email == "" {
		utils.ReturnError(ctx, err, http.StatusNotFound)
		return
	}

	if userData.Password != user["password"] {
		utils.ReturnError(ctx, errors.New("password not matched"), http.StatusUnauthorized)
		return
	}
	utils.ReturnResponse(ctx, userData, http.StatusOK)
}

func GetUser(ctx *gin.Context) {
	data, err := userService.GetUser()
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)

		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)

}
