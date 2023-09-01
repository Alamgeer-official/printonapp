package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/services"
)

var userService = services.NewUserService(repository.NewUserRepo())

func Signup(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := userService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusCreated, data)
}
func Login(ctx *gin.Context) {
	//
}

func GetUser(ctx *gin.Context) {
	data, err := userService.GetUser()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusAccepted, data)
}
