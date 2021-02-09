package apperror

import "github.com/pkg/errors"

// AppError .
type AppError struct {
	Raw       error
	ErrorCode int
	HTTPCode  int
	Message   string
}

func (e AppError) Error() string {
	if e.Raw != nil {
		return errors.Wrap(e.Raw, e.Message).Error()
	}

	return e.Message
}

// NewError .
func NewError(err error, httpCode, errCode int, message string) AppError {
	return AppError{
		Raw:       err,
		ErrorCode: errCode,
		HTTPCode:  httpCode,
		Message:   message,
	}
}
