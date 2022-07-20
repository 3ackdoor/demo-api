package middleware

import (
	"net/http"

	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/3ackdoor/go-demo-api/src/util/response"
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
	switch e := err.(type) {
	case *exception.InternalServiceException:
		response.Fail(c, exception.StatusCode(exception.InternalServiceExceptionCode), e)
	case *exception.InputValidationException:
		response.Fail(c, exception.StatusCode(exception.InputValidationExceptionCode), e)
	case *exception.ValidationException:
		response.Fail(c, exception.StatusCode(exception.ValidationExceptionCode), e)
	case *exception.ExternalServiceException:
		response.Fail(c, exception.StatusCode(exception.ExternalServiceExceptionCode), e)
	default:
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}
