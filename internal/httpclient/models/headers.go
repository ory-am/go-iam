// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Headers Headers is the jwt headers
//
// swagger:model Headers
type Headers struct {

	// extra
	Extra interface{} `json:"extra,omitempty"`
}

// Validate validates this headers
func (m *Headers) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this headers based on context it is used
func (m *Headers) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Headers) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Headers) UnmarshalBinary(b []byte) error {
	var res Headers
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
