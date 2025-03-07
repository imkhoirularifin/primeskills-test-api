package middleware

import (
	"errors"
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/pkg/exception"
	"primeskills-test-api/pkg/xlogger"

	"github.com/gin-gonic/gin"
)

// HandleError is a middleware function for handling errors in Gin.
// It logs the error and sends an appropriate JSON response based on the error type.
func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := xlogger.Logger
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			var httpException *exception.HTTPException
			if errors.As(err.Err, &httpException) {
				// If the error is an HTTPException, respond with its status code and message.
				c.JSON(httpException.StatusCode, &dto.ResponseDto{
					Message: httpException.Message,
				})
				return
			}

			// For any other errors, respond with a 500 Internal Server Error.
			c.JSON(http.StatusInternalServerError, &dto.ResponseDto{
				Message: "Something went wrong",
			})

			// Log the error.
			logger.Error().Msgf("Error: %v", err.Err)
		}
	}
}
