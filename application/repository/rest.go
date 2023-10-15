package repository

import (
	"errors"
	"net/http"
	"nu/corpus-reader/application/domain"
)

func createNewError(message string, status int) *domain.RestError {
	var error = errors.New(message)
	return &domain.RestError{
		Error:  error,
		Status: status,
	}
}

func BadRequestError(message string) *domain.RestError {
	return createNewError(message, http.StatusBadRequest)
}

func ServiceUnavailableError(message string) *domain.RestError {
	return createNewError(message, http.StatusServiceUnavailable)
}
