package errors

import (
	"fmt"
	"net/http"
)

type AError struct {
	StatusCodeVal   int    `json:"status"`
	ErrorMessageVal string `json:"error"`
	DetailsVal      string `json:"details"`
}

type APIError interface {
	error
	StatusCode() int
	ErrorMessage() string
	Details() string
}

func (e *AError) Error() string {
	return e.ErrorMessageVal
}

func (e *AError) StatusCode() int {
	return e.StatusCodeVal
}

func (e *AError) ErrorMessage() string {
	return e.ErrorMessageVal
}

func (e *AError) Details() string {
	return e.DetailsVal
}

type KeyNotFoundError struct {
	*AError
}
type KeyAlreadyExistsError struct {
	*AError
}

type InvalidReceiptError struct {
	*AError
}
type InternalServerError struct {
	*AError
}

func NewInternalServerError() *InternalServerError {
	return &InternalServerError{
		AError: &AError{StatusCodeVal: http.StatusInternalServerError,
			DetailsVal:      "Internal Server Error",
			ErrorMessageVal: "An unexpected error occurred",
		},
	}
}

func NewKeyAlreadyExistsError(id string) *KeyAlreadyExistsError {
	return &KeyAlreadyExistsError{
		AError: &AError{StatusCodeVal: http.StatusConflict,
			DetailsVal:      fmt.Sprintf("Receipt with ID: %s already exists", id),
			ErrorMessageVal: "Receipt already exists",
		},
	}
}

func NewKeyNotFoundError(id string) *KeyNotFoundError {
	return &KeyNotFoundError{
		AError: &AError{StatusCodeVal: http.StatusNotFound,
			ErrorMessageVal: "No receipt found for that ID.",
			DetailsVal:      fmt.Sprintf("ID: %s does not exist", id),
		},
	}
}

func NewInvalidReceiptError(reason string) *InvalidReceiptError {
	return &InvalidReceiptError{
		AError: &AError{StatusCodeVal: http.StatusBadRequest,
			ErrorMessageVal: "The receipt is invalid.",
			DetailsVal:      reason},
	}
}
