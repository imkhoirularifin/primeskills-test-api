package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// Logger returns a gin.HandlerFunc that logs HTTP requests using zerolog.
// It logs the method, URL, client IP, status code, and duration of each request.
// If the status code is 400 or higher, it logs the request as an error.
func Logger(logger *zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		statusCode := c.Writer.Status()
		isError := statusCode >= 400

		level := logger.Info()
		if isError {
			level = logger.Error()
		}

		level.Str("method", c.Request.Method).
			Str("url", c.Request.URL.String()).
			Str("client_ip", c.ClientIP()).
			Int("status", statusCode).
			Dur("duration", duration).
			Msg("HTTP request")
	}
}
