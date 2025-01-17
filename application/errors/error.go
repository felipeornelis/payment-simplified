package errors

import "fmt"

type AppError struct {
	Code    int
	Message string
	Err     error
}

func New(code int, err error) error {
	return AppError{
		Code:    code,
		Message: err.Error(),
		Err:     err,
	}
}

// Error implements error.
func (a AppError) Error() string {
	return fmt.Sprintf("Code: %d, Error: %v", a.Code, a.Err)
}
