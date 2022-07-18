package middleware

import (
	"bytes"
	"log"

	"github.com/gin-gonic/gin"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

/*
	no need to implement WriteString method
	source: https://github.com/gin-gonic/gin/issues/1363
*/
// func (r responseBodyWriter) WriteString(s string) (n int, err error)  {
// 	r.body.WriteString(s)
// 	return r.ResponseWriter.WriteString(s)
// }


func LogHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer}
		ctx.Writer = w
		ctx.Next()
		log.Println("Response body: " + w.body.String())

		w.body.Reset()
	}
}
