package util

import "github.com/gin-gonic/gin"

type ResponseUtil struct {
}

func ResponseSuccess(ctx *gin.Context, data any) {
	ctx.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}
