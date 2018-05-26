// Code generated by go-swagger; DO NOT EDIT.

package swagmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// RegisterResponse register response
// swagger:model registerResponse
type RegisterResponse struct {

	// uid
	UID string `json:"uid,omitempty"`

	// user ID
	UserID string `json:"userID,omitempty"`
}

// Validate validates this register response
func (m *RegisterResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RegisterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RegisterResponse) UnmarshalBinary(b []byte) error {
	var res RegisterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
