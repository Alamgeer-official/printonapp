package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/services"
	"githuh.com/printonapp/utils"
)

type ThesisCtr interface {
	CreateThesis(ctx *gin.Context)
	ReadAllTheses(ctx *gin.Context)
	ReadAllThesesByRole(ctx *gin.Context)
	GetThesisByID(ctx *gin.Context)
}

type thesisCtr struct{}

func NewThesisCtr() ThesisCtr {
	return &thesisCtr{}
}

var thesisSrv = services.NewThesisSrv(repository.NewThesisRepo())

func (tc *thesisCtr) CreateThesis(ctx *gin.Context) {
	var thesis models.Theses
	if err := ctx.ShouldBindJSON(&thesis); err != nil {
		utils.ReturnError(ctx, err, http.StatusBadRequest)
		return
	}

	if err := thesisSrv.CreateThesis(ctx, &thesis); err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}

	utils.ReturnResponse(ctx, "Thesis created successfully", http.StatusCreated)
}

func (tc *thesisCtr) ReadAllTheses(ctx *gin.Context) {
	data, err := thesisSrv.ReadAllTheses(ctx)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)
}

func (tc *thesisCtr) ReadAllThesesByRole(ctx *gin.Context) {
	// Parse pagination parameters
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(ctx.DefaultQuery("pageSize", "10"))
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	data, err := thesisSrv.ReadAllThesesByRole(ctx, page, pageSize)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)
}

func (tc *thesisCtr) GetThesisByID(ctx *gin.Context) {
	id := ctx.Param("id")
	thesisID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		utils.ReturnError(ctx, errors.New("invalid thesis ID"), http.StatusBadRequest)
		return
	}

	thesis, err := thesisSrv.GetThesisByID(ctx, thesisID)
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusNotFound)
		return
	}

	utils.ReturnResponse(ctx, thesis, http.StatusOK)
}
