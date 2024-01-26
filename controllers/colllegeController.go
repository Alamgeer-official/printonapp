package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/repository"
	"githuh.com/printonapp/services"
	"githuh.com/printonapp/utils"
)

type CollegeCtr interface {
	GetColleges(ctx *gin.Context)
}

type collegeCtr struct{}

func NewCollegeCtr() CollegeCtr {

	return &collegeCtr{}
}

var collegeSrv = services.NewCollegeRepo(repository.NewCollegeRepo())

func (cc *collegeCtr)GetColleges(ctx *gin.Context) {

	data, err := collegeSrv.GetColleges()
	if err != nil {
		utils.ReturnError(ctx, err, http.StatusInternalServerError)
		return
	}
	utils.ReturnResponse(ctx, data, http.StatusOK)

}
