package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func Zerolog(logger *zerolog.Logger) gin.HandlerFunc {
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
