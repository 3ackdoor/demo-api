package middleware

import (
	"net/http"

	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/gin-gonic/gin"
)

const (
	recovered string = "recovered"
)

func RecoveryWriterHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		MapErrorResponse(err, c)
		c.Set(recovered, true)
	}
}

// MapErrorResponse serializes the presented error as JSON into response body
func MapErrorResponse(err any, c *gin.Context) {
	var statusCode int
	switch e := err.(type) {
	case *exception.InternalServiceException:
		statusCode = exception.StatusCode(exception.InternalServiceExceptionCode)
		c.JSON(statusCode, e)
	case *exception.InputValidationException:
		statusCode = exception.StatusCode(exception.InputValidationExceptionCode)
		c.JSON(statusCode, e)
	case *exception.ValidationException:
		statusCode = exception.StatusCode(exception.ValidationExceptionCode)
		c.JSON(statusCode, e)
	case *exception.ExternalServiceException:
		statusCode = exception.StatusCode(exception.ExternalServiceExceptionCode)
		c.JSON(statusCode, e)
	default:
		statusCode = http.StatusInternalServerError
		c.AbortWithStatus(statusCode)
	}
}
