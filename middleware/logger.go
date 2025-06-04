package middleware

import (
	"github.com/gin-gonic/gin"
	"pkg/logging"
	"time"

	//import uuid
	"github.com/google/uuid"
)

const (
	HeaderRequestID    = "X-Request-Id"
	HeaderTraceID      = "X-Trace-Id"
	HeaderForwardedFor = "X-Forwarded-For"

	XRequestID = "x-request-id"
)

func Logger(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		reqID := c.GetHeader(HeaderRequestID)
		if reqID == "" {
			reqID = uuid.NewString()
			//set into context
			c.Set(XRequestID, reqID)
		}
		endpoint := c.FullPath()

		l := logging.
			DefaultLogger().
			With("request_id", reqID).
			With("endpoint", endpoint).
			With("method", c.Request.Method).
			With("host", c.Request.Host).
			With("uri", c.Request.RequestURI).
			With("user_agent", c.Request.UserAgent()).
			With("client_ip", c.ClientIP()).
			With("referer", c.Request.Referer())

		//populate logger into context
		c.Next()
		duration := time.Since(start)

		l = l.With("duration", duration)

		if c.Writer.Status() >= 200 && c.Writer.Status() <= 299 {
			l.Infof("success status code %d", c.Writer.Status())
		} else {
			l.Infof("error status code %d", c.Writer.Status())
		}
	}
}
