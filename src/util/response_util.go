package util

import "github.com/gin-gonic/gin"

func ResponseSuccess(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"message": "success",
		"data":    data,
	})
}
