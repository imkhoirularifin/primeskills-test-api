package middleware

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/internal/utilities"
	"primeskills-test-api/pkg/xlogger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Validate[V any]() gin.HandlerFunc {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return func(c *gin.Context) {
		logger := xlogger.Logger

		var v V
		if err := c.ShouldBind(&v); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request",
			})
			logger.Error().Msgf("Error: %v", err)
			c.Abort()
			return
		}

		if err := validate.Struct(v); err != nil {
			var errors []string
			for _, err := range err.(validator.ValidationErrors) {
				errMsg := utilities.GetValidationErrorMessage(err)
				errors = append(errors, errMsg)
			}
			c.JSON(http.StatusBadRequest, &dto.ResponseDto{
				Errors:  errors,
				Message: http.StatusText(http.StatusBadRequest),
			})
			c.Abort()
			return
		}

		c.Set("parser", &v)
		c.Next()
	}
}
