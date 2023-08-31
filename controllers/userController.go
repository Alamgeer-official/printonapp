package controllers

import "github.com/gin-gonic/gin"

func Signup(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Signup",
	})
}
func Login(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Login",
	})
}
