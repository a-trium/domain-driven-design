package exception

import (
	"time"
)

const CodeBadRequest = 400
const CodeNotFound = 404
const CodeInternalServer = 500

type Exception interface {
	Cause() error
	StatusCode() int
	Timestamp() *time.Time

	IsBadRequestError() bool
	IsNotFoundError() bool
	IsInternalServerError() bool
}

type appException struct {
	cause      error
	statusCode int
	timestamp  time.Time
}

func (a *appException) Cause() error {
	return a.cause
}

func (a *appException) StatusCode() int {
	return a.statusCode
}

func (a *appException) Timestamp() *time.Time {
	return &a.timestamp
}

func (a *appException) IsBadRequestError() bool {
	return a.statusCode == CodeBadRequest
}

func NewBadRequestError(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeBadRequest,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsNotFoundError() bool {
	return a.statusCode == CodeNotFound
}

func NewNotFoundError(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeNotFound,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsInternalServerError() bool {
	return a.statusCode == CodeInternalServer
}

func NewInternalServerError(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeInternalServer,
		timestamp:  time.Now(),
	}
}
