package middleware

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/3ackdoor/go-demo-api/src/exception"
	"github.com/gin-gonic/gin"
)

func CustomRecoveryFunc() gin.RecoveryFunc {
	return func(c *gin.Context, errMsg any) {
		log.Printf("err: %v", errMsg)
		log.Printf("custom recovery func")

		errVal, ok := errMsg.(string)
		if !ok {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		var validationException exception.ValidationException
		if err := json.Unmarshal([]byte(errVal), &validationException); err == nil && validationException.Code == exception.ValidationExceptionName {
			log.Printf("validationException: %v", validationException)
		}

		var externalServiceException exception.ExternalServiceException
		if err := json.Unmarshal([]byte(errVal), &externalServiceException); err == nil && externalServiceException.Code == exception.ExternalServiceExceptionName {
			log.Printf("externalServiceException: %v", externalServiceException)
		}

	}
}