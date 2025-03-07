package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

// GetValidationErrorMessage returns a user-friendly validation error message
// based on the provided FieldError. It uses a predefined set of validation
// messages and formats them with the field name and any additional parameters.
//
// Parameters:
//   - e: validator.FieldError - the validation error to process
//
// Returns:
//   - string: the formatted validation error message
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

// validationMessages is a map that associates validation tags with their
// corresponding user-friendly error messages. The messages can include
// placeholders for additional parameters.
var validationMessages = map[string]string{
	"required": "is required",
	"email":    "must be a valid email address",
	"min":      "must be at least %s characters",
	"max":      "must be at most %s characters",
	"uuid4":    "must be a valid UUID v4",
}
