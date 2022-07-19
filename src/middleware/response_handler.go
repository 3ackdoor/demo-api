package middleware

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	noWritten = -1
)

// ResponseWriterHandler is the implementation of http.ResponseWriter
func ResponseWriterHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		/* Before Request */
		w := NewResponseWriter(c.Writer)
		c.Writer = w

		c.Next()

		/* After Request */
		_, exists := c.Get("recovered")
		log.Printf("recovered: %v ", exists)
		if !exists && strings.Contains(w.Header().Get("Content-Type"), gin.MIMEJSON) {
			data := w.Body.Bytes()
			w.Body.Reset()

			resp := NewJsonResponse(data)
			body, err := json.Marshal(resp)
			if err != nil {
				log.Printf("Something went wrong while encoding JSON: %v", err)
				return
			}

			w.Body.Write(body)
		} else {
			body := w.Body.Bytes()
			w.Body.Reset()

			w.Body.Write(body)
		}

		w.Flush()
	}
}

type jsonResponse struct {
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}

func NewJsonResponse(data any) *jsonResponse {
	return &jsonResponse{http.StatusText(http.StatusOK), data}
}

func (j *jsonResponse) MarshalJSON() ([]byte, error) {
	msg, err := json.Marshal(j.Message)
	if err != nil {
		return nil, err
	}
	b := fmt.Sprintf("{\"message\":%s,\"data\":%s}", msg, j.Data)
	return []byte(b), nil
}

type responseWriter struct {
	Response gin.ResponseWriter // the actual ResponseWriter to flush to
	status   int                // the HTTP response code from WriteHeader
	Body     *bytes.Buffer      // the response content body
	Flushed  bool
}

func NewResponseWriter(w gin.ResponseWriter) *responseWriter {
	return &responseWriter{
		Response: w, status: http.StatusOK, Body: &bytes.Buffer{},
	}
}

func (w *responseWriter) Header() http.Header {
	return w.Response.Header() // use the actual response header
}

func (w *responseWriter) Write(buf []byte) (int, error) {
	w.Body.Write(buf)
	return len(buf), nil
}

func (w *responseWriter) WriteString(s string) (n int, err error) {
	n, err = w.Write([]byte(s))
	return
}

func (w *responseWriter) Written() bool {
	return w.Body.Len() != noWritten
}

func (w *responseWriter) WriteHeader(status int) {
	w.status = status
}

func (w *responseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.Response.WriteHeader(w.status)
	}
}

func (w *responseWriter) Status() int {
	return w.status
}

func (w *responseWriter) Size() int {
	return w.Body.Len()
}

func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.Response.(http.Hijacker).Hijack()
}

func (w *responseWriter) CloseNotify() <-chan bool {
	return w.Response.CloseNotify()
}

func (w *responseWriter) Flush() {
	if w.Flushed {
		return
	}
	w.Response.WriteHeader(w.status)
	if w.Body.Len() > 0 {
		_, err := w.Response.Write(w.Body.Bytes())
		if err != nil {
			panic(err)
		}
		w.Body.Reset()
	}
	w.Flushed = true
}

func (w *responseWriter) Pusher() (pusher http.Pusher) {
	if pusher, ok := w.Response.(http.Pusher); ok {
		return pusher
	}
	return nil
}
