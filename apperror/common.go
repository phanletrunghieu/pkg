package apperror

import (
	"net/http"
)

// ErrUnauthorize .
func ErrUnauthorize(err error) AppError {
	return AppError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: 100010,
		Message:   "Unauthorized!",
	}
}

// ErrCommitTransaction .
func ErrCommitTransaction(err error) AppError {
	return AppError{
		Raw:       err,
		HTTPCode:  http.StatusUnauthorized,
		ErrorCode: 100020,
		Message:   "Fail to commit transaction",
	}
}
