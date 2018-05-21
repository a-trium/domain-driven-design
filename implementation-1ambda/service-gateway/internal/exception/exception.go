package exception

import (
	"time"

	dto "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
)

const CodeBadRequest = 400
const CodeUnauthorized = 401
const CodeForbidden = 403
const CodeNotFound = 404
const CodeInternalServer = 500

type Exception interface {
	Cause() error
	StatusCode() int
	Timestamp() *time.Time
	ToSwaggerError() *dto.Error

	IsBadRequestException() bool
	IsUnauthorizedException() bool
	IsForbiddenException() bool
	IsNotFoundException() bool
	IsInternalServerException() bool
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

func (a *appException) ToSwaggerError() *dto.Error {
	errorType := ""

	switch status := a.statusCode; status {
	case CodeBadRequest:
		errorType = dto.ErrorTypeBadRequest

	case CodeForbidden:
		errorType = dto.ErrorTypeForbidden

	case CodeUnauthorized:
		errorType = dto.ErrorTypeUnauthorized

	case CodeNotFound:
		errorType = dto.ErrorTypeNotFound

	default:
		errorType = dto.ErrorTypeInternalServer
	}

	return &dto.Error{
		Code:      int64(a.statusCode),
		Message:   a.cause.Error(),
		Timestamp: a.timestamp.UTC().String(),
		Type:      errorType,
	}
}

func (a *appException) IsBadRequestException() bool {
	return a.statusCode == CodeBadRequest
}

func NewBadRequestException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeBadRequest,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsUnauthorizedException() bool {
	return a.statusCode == CodeUnauthorized
}

func NewUnauthorizedException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeUnauthorized,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsForbiddenException() bool {
	return a.statusCode == CodeForbidden
}

func NewForbiddenException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeForbidden,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsNotFoundException() bool {
	return a.statusCode == CodeNotFound
}

func NewNotFoundException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeNotFound,
		timestamp:  time.Now(),
	}
}

func (a *appException) IsInternalServerException() bool {
	return a.statusCode == CodeInternalServer
}

func NewInternalServerException(err error) Exception {
	return &appException{
		cause:      err,
		statusCode: CodeInternalServer,
		timestamp:  time.Now(),
	}
}
