package middleware

import (
	"errors"
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/pkg/exception"
	"primeskills-test-api/pkg/xlogger"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := xlogger.Logger
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			var httpException *exception.HTTPException
			if errors.As(err.Err, &httpException) {
				c.JSON(httpException.StatusCode, &dto.ResponseDto{
					Message: httpException.Message,
				})
				return
			}

			c.JSON(http.StatusInternalServerError, &dto.ResponseDto{
				Message: "Something went wrong",
			})

			logger.Error().Msgf("Error: %v", err.Err)
		}
	}
}
