package model

import "net/http"

type Error struct {
	status  int
	code    string
	message string
}

var (
	IinEmpty      = New(http.StatusBadRequest, "bad_request", "Iin is empty")
	NameEmpty     = New(http.StatusBadRequest, "bad_request", "Name is empty")
	ErrIinInvalid = New(http.StatusBadRequest, "bad_request", "IIN is invalid")
)

func New(status int, code, message string) *Error {
	return &Error{
		status:  status,
		code:    code,
		message: message,
	}
}

func (e Error) Error() string {
	return e.message
}
