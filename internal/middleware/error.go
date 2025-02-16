package middleware

import (
	"net/http"
	"primeskills-test-api/pkg/xerrors"
	"primeskills-test-api/pkg/xlogger"

	"github.com/gin-gonic/gin"
)

func HandleError() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := xlogger.Logger
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()

			if customError, ok := err.Err.(*xerrors.CustomError); ok {
				c.JSON(customError.Status, gin.H{
					"message": customError.Message,
				})
				return
			}

			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Something went wrong",
			})
			logger.Error().Msgf("Error: %v", err.Err)
		}
	}
}
