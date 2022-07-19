package middleware

import (
	"net/http"

	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		MapErrorResponse(err, c)
	}
}

func MapErrorResponse(error any, c *gin.Context) {
	switch err := error.(type) {
	case *exception.InternalServiceException:
		c.JSON(exception.StatusCode(exception.InternalServiceExceptionCode), err)
	case *exception.InputValidationException:
		c.JSON(exception.StatusCode(exception.InputValidationExceptionCode), err)
	case *exception.ValidationException:
		c.JSON(exception.StatusCode(exception.ValidationExceptionCode), err)
	case *exception.ExternalServiceException:
		c.JSON(exception.StatusCode(exception.ExternalServiceExceptionCode), err)
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
