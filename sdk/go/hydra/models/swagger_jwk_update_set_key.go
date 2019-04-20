// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SwaggerJwkUpdateSetKey SwaggerJwkUpdateSetKey swagger jwk update set key
// swagger:model SwaggerJwkUpdateSetKey
type SwaggerJwkUpdateSetKey struct {

	// body
	Body *SwaggerJSONWebKey `json:"Body,omitempty"`

	// The kid of the desired key
	// in: path
	// Required: true
	KID *string `json:"kid"`

	// The set
	// in: path
	// Required: true
	Set *string `json:"set"`
}

// Validate validates this swagger jwk update set key
func (m *SwaggerJwkUpdateSetKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwaggerJwkUpdateSetKey) validateBody(formats strfmt.Registry) error {

	if swag.IsZero(m.Body) { // not required
		return nil
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Body")
			}
			return err
		}
	}

	return nil
}

func (m *SwaggerJwkUpdateSetKey) validateKID(formats strfmt.Registry) error {

	if err := validate.Required("kid", "body", m.KID); err != nil {
		return err
	}

	return nil
}

func (m *SwaggerJwkUpdateSetKey) validateSet(formats strfmt.Registry) error {

	if err := validate.Required("set", "body", m.Set); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwaggerJwkUpdateSetKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwaggerJwkUpdateSetKey) UnmarshalBinary(b []byte) error {
	var res SwaggerJwkUpdateSetKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
