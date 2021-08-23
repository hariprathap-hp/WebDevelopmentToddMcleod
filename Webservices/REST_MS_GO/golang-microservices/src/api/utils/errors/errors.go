package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Rstatus  int    `json:"status"`
	Rmessage string `json:"message"`
	Rerror   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.Rstatus
}
func (e *apiError) Message() string {
	return e.Rmessage
}
func (e *apiError) Error() string {
	return e.Rerror
}

func NewNotFoundError(message string) ApiError {
	return &apiError{
		Rstatus:  http.StatusNotFound,
		Rmessage: message,
	}
}

func NewInternalServersError(message string) ApiError {
	return &apiError{
		Rstatus:  http.StatusInternalServerError,
		Rmessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		Rstatus:  http.StatusBadRequest,
		Rmessage: message,
	}
}

func NewApiError(statuscode int, message string) ApiError {
	return &apiError{
		Rstatus:  statuscode,
		Rmessage: message,
	}
}
