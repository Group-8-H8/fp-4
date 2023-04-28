package errs

import "net/http"

type errs struct {
	ErrMessage string
	ErrStatus  int
	ErrError   string
}

type Errs interface {
	Message() string
	Status() int
	Error() string
}

func (err *errs) Message() string {
	return err.ErrMessage
}

func (err *errs) Status() int {
	return err.ErrStatus
}

func (err *errs) Error() string {
	return err.ErrError
}

func NewInternalServerError(message string) Errs {
	return &errs{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError: "INTERNAL_SERVER_ERROR",
	}
}

func NewBadRequestError(message string) Errs {
	return &errs{
		ErrMessage: message,
		ErrStatus: http.StatusBadRequest,
		ErrError: "BAD_REQUEST",
	}
}

func NewUnprocessableEntityError(message string) Errs {
	return &errs{
		ErrMessage: message,
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: "INVALID_REQUEST_BODY",
	}
}

func NewNotFoundError(message string) Errs {
	return &errs{
		ErrMessage: message,
		ErrStatus: http.StatusNotFound,
		ErrError: "NOT_FOUND",
	}
}

func NewUnauthenticatedError(message string) Errs {
	return &errs{
		ErrMessage: message,
		ErrStatus: http.StatusUnauthorized,
		ErrError: "NOT_AUTHENTICATED",
	}
}

func NewUnauthorizedError(message string) Errs {
	return &errs{
		ErrMessage: message,
		ErrStatus: http.StatusForbidden,
		ErrError: "NOT_AUTHORIZED",
	}
}