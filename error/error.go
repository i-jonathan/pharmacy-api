package appError

import "net/http"

type AppError struct {
	Status  int
	Message string
}

var (
	BadRequest   = AppError{http.StatusBadRequest, "Bad request body received"}
	ServerError  = AppError{http.StatusInternalServerError, "An error occurred while processing that request"}
	Forbidden    = AppError{http.StatusForbidden, "Forbidden"}
	NotFound     = AppError{http.StatusNotFound, "The requested resource was not found"}
	Unauthorized = AppError{http.StatusUnauthorized, "Unauthorized"}
)

func NewError(status int, msg string) AppError {
	return AppError{status, msg}
}

func (err AppError) Error() string {
	return err.Message
}

func (err AppError) Response() map[string]string {
	return map[string]string{"message": err.Message}
}
