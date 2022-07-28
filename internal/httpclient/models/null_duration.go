// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NullDuration NullDuration represents a nullable JSON and SQL compatible time.Duration.
//
// TODO delete this type and replace it with ory/x/sqlxx/NullDuration when applying the custom client token TTL patch to Hydra 2.x
//
// swagger:model nullDuration
type NullDuration struct {

	// duration
	Duration Duration `json:"Duration,omitempty"`

	// valid
	Valid bool `json:"Valid,omitempty"`
}

// Validate validates this null duration
func (m *NullDuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NullDuration) validateDuration(formats strfmt.Registry) error {
	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if err := m.Duration.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Duration")
		}
		return err
	}

	return nil
}

// ContextValidate validate this null duration based on the context it is used
func (m *NullDuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDuration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NullDuration) contextValidateDuration(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Duration.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Duration")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NullDuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NullDuration) UnmarshalBinary(b []byte) error {
	var res NullDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
