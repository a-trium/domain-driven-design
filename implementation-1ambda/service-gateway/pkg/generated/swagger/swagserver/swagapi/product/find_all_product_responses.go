// Code generated by go-swagger; DO NOT EDIT.

package product

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	swagmodel "github.com/a-trium/domain-driven-design/implementation-1ambda/service-gateway/pkg/generated/swagger/swagmodel"
)

// FindAllProductOKCode is the HTTP code returned for type FindAllProductOK
const FindAllProductOKCode int = 200

/*FindAllProductOK OK

swagger:response findAllProductOK
*/
type FindAllProductOK struct {

	/*
	  In: Body
	*/
	Payload *swagmodel.FindAllProductOKBody `json:"body,omitempty"`
}

// NewFindAllProductOK creates FindAllProductOK with default headers values
func NewFindAllProductOK() *FindAllProductOK {

	return &FindAllProductOK{}
}

// WithPayload adds the payload to the find all product o k response
func (o *FindAllProductOK) WithPayload(payload *swagmodel.FindAllProductOKBody) *FindAllProductOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find all product o k response
func (o *FindAllProductOK) SetPayload(payload *swagmodel.FindAllProductOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindAllProductOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindAllProductDefault error

swagger:response findAllProductDefault
*/
type FindAllProductDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *swagmodel.Exception `json:"body,omitempty"`
}

// NewFindAllProductDefault creates FindAllProductDefault with default headers values
func NewFindAllProductDefault(code int) *FindAllProductDefault {
	if code <= 0 {
		code = 500
	}

	return &FindAllProductDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find all product default response
func (o *FindAllProductDefault) WithStatusCode(code int) *FindAllProductDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find all product default response
func (o *FindAllProductDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find all product default response
func (o *FindAllProductDefault) WithPayload(payload *swagmodel.Exception) *FindAllProductDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find all product default response
func (o *FindAllProductDefault) SetPayload(payload *swagmodel.Exception) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindAllProductDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
