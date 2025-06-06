package middleware

import (
	"github.com/gin-gonic/gin"
	"pkg/logging"
	"time"

	"github.com/google/uuid"
)

const (
	HeaderRequestID    = "X-Request-Id"
	HeaderTraceID      = "X-Trace-Id"
	HeaderForwardedFor = "X-Forwarded-For"

	XRequestID = "x-request-id"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqID := c.GetHeader(HeaderRequestID)

		if reqID == "" {
			reqID = uuid.NewString()
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

		c.Next()

		latency := time.Since(start).Milliseconds()
		status := c.Writer.Status()
		status = 502

		l = l.
			With("latency_ms", latency).
			With("status", status)

		switch {
		case status >= 200 && status < 400:
			l.Infof("success")
		case status >= 400 && status < 500:
			l.Warn("client error")
		case status >= 500 && status < 600:
			l.Error("server error")
		default:
			l.Warn("unknown error")
		}
	}
}
