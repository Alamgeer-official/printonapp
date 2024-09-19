package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"githuh.com/printonapp/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//extract token
		tokenString, err := utils.ExtractToken(ctx)
		if err != nil {
			utils.ReturnError(ctx, err, http.StatusUnauthorized)
			ctx.Abort()
			return
		}

		//verify token
		var token *jwt.Token
		token, err = utils.VerifyToken(tokenString)
		if err != nil {
			utils.ReturnError(ctx, err, http.StatusUnauthorized)
			ctx.Abort()
		}

		// Get meta data from token
		tokenMetaData, err := utils.GetMetaDataFromToken(token)
		if err != nil {
			utils.ReturnError(ctx, err, http.StatusUnauthorized)
			ctx.Abort()
		}

		// Add metadata to the Gin context
		ctx.Set("user", tokenMetaData.User)

		ctx.Next()
	}

}
