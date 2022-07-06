package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomRecoveryFunc() gin.RecoveryFunc {
	return func(c *gin.Context, err any) {
		log.Printf("err: %v", err)
		log.Printf("custom recovery func")

		// json.Unmarshal(data []byte, v any)


		if err, ok := err.(string); ok {
			c.String(http.StatusInternalServerError, fmt.Sprintf("error: %s", err))
		}
		c.AbortWithStatus(http.StatusInternalServerError)

	}
}