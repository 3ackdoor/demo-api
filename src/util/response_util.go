package util

import "github.com/gin-gonic/gin"

func ResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}
