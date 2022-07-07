package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/3ackdoor/go-demo-api/src/dto"
	"github.com/gin-gonic/gin"
)

func CustomRecoveryFunc() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		log.Printf("err: %v", err)
		log.Printf("custom recovery func")

		tValidationError := &dto.ValidationError{
			Key: "kkkkkkkkkkkkkkkkkkkkk",
			Value: "vvvvvvvvvvvvvvvvv",
			Description: "dddddddddddddddddddd",
		}
		v, err := json.Marshal(tValidationError)
		if err != nil {
			log.Panic(err)
		}

		log.Printf("value: %v", string(v))

		var data dto.ValidationError
		log.Printf("init: %v", data)
		_ = json.Unmarshal(v, &data);

		log.Printf("new value: %v", data)

		if err, ok := err.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)

	}
}