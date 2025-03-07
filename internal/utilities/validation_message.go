package utilities

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetValidationErrorMessage(e validator.FieldError) string {
	tag := e.Tag()
	field := e.Field()
	param := e.Param()

	if message, ok := validationMessages[tag]; ok {
		if param != "" {
			return field + " " + fmt.Sprintf(message, param)
		}
		return field + " " + message
	}

	return field + " is invalid"
}

var validationMessages = map[string]string{
	"required": "is required",
	"email":    "must be a valid email address",
	"min":      "must be at least %s characters",
	"max":      "must be at most %s characters",
	"uuid4":    "must be a valid UUID v4",
}
