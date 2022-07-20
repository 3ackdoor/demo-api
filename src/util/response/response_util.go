package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func Fail(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}
