package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/models"
)

var (
	SECRET = os.Getenv("SECRET")
)

func CreateJWToken(user *models.User) (string, error) {
	//omit confidential info
	var expiryTime int64
	if user.Role == "ADMIN" {
		expiryTime = time.Now().Add(time.Hour * 24).Unix()
	} else {
		expiryTime = time.Now().Add(time.Hour * 4).Unix()
	}
	user.Password = ""
	user.Phone = ""
	atClaim := models.Claim{
		Id:   user.ID,
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
			Issuer:    "printonapp",
		},
	}

	tokenWithClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaim)
	token, err := tokenWithClaim.SignedString([]byte(SECRET))
	if err != nil {
		return "", nil
	}

	return token, nil

}

func ExtractToken(ctx *gin.Context) (string, error) {
	// extract token
	tokenString := ctx.GetHeader("Authorization")
	bearerToken := strings.Split(tokenString, " ")
	if len(bearerToken) != 2 {
		return "", errors.New("inappropriate token")
	}
	return bearerToken[1], nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	// Parse and validate the token in one step
	token, err := jwt.ParseWithClaims(tokenString, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signin method: %v", token.Header["alg"])
		}
		return []byte(SECRET), nil
	})
	if err != nil {
		return nil, err
	}

	// Check if the token is valid (this is redundant with Parse, but doesn't hurt)
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}

	return token, nil
}

func GetMetaDataFromToken(token *jwt.Token) (*models.Claim, error) {

	// Type-assert the claims to your custom Claim type
	claims, ok := token.Claims.(*models.Claim)
	if !ok {
		return nil, errors.New("error in getting metadata")
	}

	return claims, nil
}

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

func GetUserDataFromContext(ctx *gin.Context) *models.User {
	var user *models.User
	value, exist := ctx.Get("user")
	if !exist {
		return user
	}
	user = value.(*models.User)
	return user

}

func CalculatePagination(totalCount, pageSize, page int64, data interface{}) *models.Pagination {
	totalPage := totalCount / pageSize
	if totalCount%pageSize != 0 {
		totalPage++
	}
	if totalPage == 0 {
		totalPage = 1
	}
	return &models.Pagination{
		TotalCount: totalCount,
		TotalPage:  int(totalPage),
		Data:       data,
	}
}
