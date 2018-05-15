package exception

import (
	"time"
)

const CodeBadRequest = 400
const CodeNotFound = 404
const CodeInternalServer = 500

type Error struct {
	Err       error
	Code      int
	Timestamp time.Time
}

func (a *Error) IsBadRequestError() bool {
	return a.Code == CodeBadRequest
}

func NewBadRequestError(err error) *Error {
	return &Error{
		Err:       err,
		Code:      CodeBadRequest,
		Timestamp: time.Now(),
	}
}

func (a *Error) IsNotFoundError() bool {
	return a.Code == CodeNotFound
}

func NewNotFoundError(err error) *Error {
	return &Error{
		Err:       err,
		Code:      CodeNotFound,
		Timestamp: time.Now(),
	}
}

func (a *Error) IsInternalServerError() bool {
	return a.Code == CodeInternalServer
}

func NewInternalServerError(err error) *Error {
	return &Error{
		Err:       err,
		Code:      CodeInternalServer,
		Timestamp: time.Now(),
	}
}
