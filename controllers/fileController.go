package controllers

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/utils"
	awssdk "githuh.com/printonapp/utils/aws_sdk"
)

type FileController struct{}

func NewFileController() *FileController {
	return &FileController{}
}

func (fc *FileController) UploadPDF(c *gin.Context) {
	// user := utils.GetUserDataFromContext(c)
	// if !user.IsUser() {
	// 	utils.ReturnError(c, errors.New("unauthorize User"), http.StatusUnauthorized)
	// }
	file, header, err := c.Request.FormFile("pdf")
	if err != nil {
		utils.ReturnError(c, errors.New("no PDF file uploaded"), http.StatusBadRequest)

		return
	}
	defer file.Close()

	// Validate file size
	const maxFileSize = 10 << 20 // 10 MB
	if header.Size > maxFileSize {
		utils.ReturnError(c, errors.New("file size exceeds 10 MB limit"), http.StatusBadRequest)
		return
	}

	// Check file extension
	fileName := header.Filename
	if !strings.HasSuffix(strings.ToLower(fileName), ".pdf") {
		utils.ReturnError(c, errors.New("PDF files are allowed"), http.StatusBadRequest)
		return
	}
	path,err:=awssdk.SaveFileS3(file,header)
	if err!= nil {
        utils.ReturnError(c, err, http.StatusInternalServerError)
        return
    }

	randomPath := path

	utils.ReturnResponse(c, gin.H{"path": randomPath}, http.StatusOK)
}
