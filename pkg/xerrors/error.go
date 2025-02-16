package xerrors

type CustomError struct {
	Status  int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func Throw(status int, message string) *CustomError {
	return &CustomError{
		Status:  status,
		Message: message,
	}
}
