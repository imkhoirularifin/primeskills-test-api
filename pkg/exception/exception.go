package exception

import (
	"net/http"
)

// HTTPException represents an HTTP error with a status code and message.
type HTTPException struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// NewHTTPException creates a new HTTPException with the given status code and message.
// If the message is empty, it uses the default HTTP status text for the given status code.
func NewHTTPException(statusCode int, message string) *HTTPException {
	var msg string
	if message == "" {
		msg = http.StatusText(statusCode)
	}

	return &HTTPException{
		StatusCode: statusCode,
		Message:    msg,
	}
}

func (e *HTTPException) Error() string {
	return e.Message
}

func BadRequest(message string) *HTTPException {
	return NewHTTPException(http.StatusBadRequest, message)
}

func Unauthorized(message string) *HTTPException {
	return NewHTTPException(http.StatusUnauthorized, message)
}

func Forbidden(message string) *HTTPException {
	return NewHTTPException(http.StatusForbidden, message)
}

func NotFound(message string) *HTTPException {
	return NewHTTPException(http.StatusNotFound, message)
}

func InternalServerError(message string) *HTTPException {
	return NewHTTPException(http.StatusInternalServerError, message)
}

func NotImplemented(message string) *HTTPException {
	return NewHTTPException(http.StatusNotImplemented, message)
}

func ServiceUnavailable(message string) *HTTPException {
	return NewHTTPException(http.StatusServiceUnavailable, message)
}

func GatewayTimeout(message string) *HTTPException {
	return NewHTTPException(http.StatusGatewayTimeout, message)
}

func Conflict(message string) *HTTPException {
	return NewHTTPException(http.StatusConflict, message)
}

func Custom(statusCode int, message string) *HTTPException {
	return NewHTTPException(statusCode, message)
}
