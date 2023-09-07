package utils

import (
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
)

func ReturnError(ctx *gin.Context, err error, statusCode int) {
	resp := models.Response{
		Status: statusCode,
		Error:  err.Error(),
		Data:   nil,
	}
	ctx.JSON(statusCode, resp)
}
func ReturnResponse(ctx *gin.Context, data any, statusCode int) {
	resp := models.Response{
		Status: statusCode,
		Error:  "",
		Data:   data,
	}
	ctx.JSON(statusCode, resp)
}
