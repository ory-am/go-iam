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

// RefreshTokenHookRequest RefreshTokenHookRequest is the request body sent to the refresh token hook.
//
// swagger:model refreshTokenHookRequest
type RefreshTokenHookRequest struct {

	// ClientID is the identifier of the OAuth 2.0 client.
	ClientID string `json:"client_id,omitempty"`

	// GrantedAudience is the list of audiences granted to the OAuth 2.0 client.
	GrantedAudience []string `json:"granted_audience"`

	// GrantedScopes is the list of scopes granted to the OAuth 2.0 client.
	GrantedScopes []string `json:"granted_scopes"`

	// requester
	Requester *OAuth2AccessRequest `json:"requester,omitempty"`

	// session
	Session *Session `json:"session,omitempty"`

	// Subject is the identifier of the authenticated end-user.
	Subject string `json:"subject,omitempty"`
}

// Validate validates this refresh token hook request
func (m *RefreshTokenHookRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequester(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSession(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RefreshTokenHookRequest) validateRequester(formats strfmt.Registry) error {
	if swag.IsZero(m.Requester) { // not required
		return nil
	}

	if m.Requester != nil {
		if err := m.Requester.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requester")
			}
			return err
		}
	}

	return nil
}

func (m *RefreshTokenHookRequest) validateSession(formats strfmt.Registry) error {
	if swag.IsZero(m.Session) { // not required
		return nil
	}

	if m.Session != nil {
		if err := m.Session.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this refresh token hook request based on the context it is used
func (m *RefreshTokenHookRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequester(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSession(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RefreshTokenHookRequest) contextValidateRequester(ctx context.Context, formats strfmt.Registry) error {

	if m.Requester != nil {
		if err := m.Requester.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("requester")
			}
			return err
		}
	}

	return nil
}

func (m *RefreshTokenHookRequest) contextValidateSession(ctx context.Context, formats strfmt.Registry) error {

	if m.Session != nil {
		if err := m.Session.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("session")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RefreshTokenHookRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RefreshTokenHookRequest) UnmarshalBinary(b []byte) error {
	var res RefreshTokenHookRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
