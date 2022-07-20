package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
}

func (r *Response) Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

func (r *Response) Fail(c *gin.Context, statusCode int, data any) {
	c.JSON(statusCode, data)
}
