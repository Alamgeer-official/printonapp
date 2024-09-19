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

// Create a new instance of the user service
var userService = services.NewUserService(repository.NewUserRepo())

// Signup godoc
// @Summary Create a new user
// @Description Signup a new user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param user body models.User true "User data"
// @Success 201 {object} models.User
// @Failure 500 {object} models.Response
// @Router /signup [post]
func Signup(ctx *gin.Context) {
    var user models.User
    err := ctx.ShouldBind(&user)
    if err != nil {
        utils.ReturnError(ctx, err, http.StatusInternalServerError)
        return
    }

    data, err := userService.CreateUser(user)
    if err != nil {
        utils.ReturnError(ctx, err, http.StatusInternalServerError)
        return
    }

    utils.ReturnResponse(ctx, data, http.StatusCreated)
}

// Login godoc
// @Summary User login
// @Description Login a user
// @Tags Users
// @Accept  json
// @Produce  json
// @Param login body map[string]string true "Login credentials"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Response
// @Router /login [post]
func Login(ctx *gin.Context) {
    var user map[string]string
    err := ctx.ShouldBind(&user)
    if err != nil {
        utils.ReturnError(ctx, err, http.StatusBadRequest)
        return
    }
    data, err := userService.Login(user)
    if err != nil {
        utils.ReturnError(ctx, err, http.StatusBadRequest)
        return
    }
    utils.ReturnResponse(ctx, data, http.StatusOK)
}

// GetUsers godoc
// @Summary Get all users
// @Description Retrieve all users
// @Tags Users
// @Produce  json
// @Success 200 {array} models.User
// @Failure 400 {object} models.Response
// @Router /users [get]
func GetUsers(ctx *gin.Context) {
    data, err := userService.GetUsers(ctx)
    if err != nil {
        utils.ReturnError(ctx, err, http.StatusBadRequest)
        return
    }
    utils.ReturnResponse(ctx, data, http.StatusOK)
}

// GetUserById godoc
// @Summary Get user by ID
// @Description Retrieve a user by ID
// @Tags Users
// @Accept  json
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} models.User
// @Failure 400 {object} models.Response
// @Router /users/{id} [get]
func GetUserById(ctx *gin.Context) {
    data, err := userService.GetUsers(ctx)
    if err != nil {
        utils.ReturnError(ctx, err, http.StatusBadRequest)
        return
    }
    utils.ReturnResponse(ctx, data, http.StatusOK)
}

// Test godoc
// @Summary Test endpoint
// @Description A simple endpoint to check server status
// @Tags General
// @Produce  json
// @Success 200 {string} string "server is running"
// @Router / [get]
func Test(ctx *gin.Context) {
    utils.ReturnResponse(ctx, "server is running", http.StatusOK)
}

// IsEmailExists godoc
// @Summary Check if email exists
// @Description Check if an email is already registered
// @Tags Users
// @Produce  json
// @Param email query string true "Email address"
// @Success 200 {boolean} boolean
// @Failure 400 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /email-exists [get]
func IsEmailExists(ctx *gin.Context) {
    email := ctx.Query("email")
    if email == "" {
        utils.ReturnError(ctx, errors.New("Email parameter is required"), http.StatusBadRequest)
        return
    }

    exists, err := userService.IsEmailExists(email)
    if err != nil {
        utils.ReturnError(ctx, errors.New("Failed to check email existence"), http.StatusInternalServerError)
        return
    }

    utils.ReturnResponse(ctx, exists, http.StatusOK)
}
