package controllers

import "github.com/gin-gonic/gin"

func UserTest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "User Its working",
	})
}
func UserTest2(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "User2 Its working",
	})
}
