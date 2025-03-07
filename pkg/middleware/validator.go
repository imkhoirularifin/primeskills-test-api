package middleware

import (
	"net/http"
	"primeskills-test-api/internal/domain/dto"
	"primeskills-test-api/pkg/utils"
	"primeskills-test-api/pkg/xlogger"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Validate is a middleware function that validates the request body against a given struct type V.
// It uses the gin context to bind the request body to the struct and validates it using the go-playground/validator package.
// If validation fails, it responds with a 400 Bad Request status and an error message.
// If validation succeeds, it sets the validated struct in the context and proceeds to the next handler.
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
				errMsg := utils.GetValidationErrorMessage(err)
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
