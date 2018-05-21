// Code generated by go-swagger; DO NOT EDIT.

package swagmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Empty empty
// swagger:model empty
type Empty struct {

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this empty
func (m *Empty) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Empty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Empty) UnmarshalBinary(b []byte) error {
	var res Empty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
