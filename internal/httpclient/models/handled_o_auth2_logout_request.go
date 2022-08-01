// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HandledOAuth2LogoutRequest The response payload sent when there is an attempt to access an already handled logout request.
//
// swagger:model handledOAuth2LogoutRequest
type HandledOAuth2LogoutRequest struct {

	// Original request URL to which you should redirect the user if request was already handled.
	// Required: true
	RedirectTo *string `json:"redirect_to"`
}

// Validate validates this handled o auth2 logout request
func (m *HandledOAuth2LogoutRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRedirectTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HandledOAuth2LogoutRequest) validateRedirectTo(formats strfmt.Registry) error {

	if err := validate.Required("redirect_to", "body", m.RedirectTo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this handled o auth2 logout request based on context it is used
func (m *HandledOAuth2LogoutRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *HandledOAuth2LogoutRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HandledOAuth2LogoutRequest) UnmarshalBinary(b []byte) error {
	var res HandledOAuth2LogoutRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
