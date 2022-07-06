// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
	"github.com/ory/hydra/internal/httpclient/models"
)

// NewAdminPatchOAuth2ClientParams creates a new AdminPatchOAuth2ClientParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminPatchOAuth2ClientParams() *AdminPatchOAuth2ClientParams {
	return &AdminPatchOAuth2ClientParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminPatchOAuth2ClientParamsWithTimeout creates a new AdminPatchOAuth2ClientParams object
// with the ability to set a timeout on a request.
func NewAdminPatchOAuth2ClientParamsWithTimeout(timeout time.Duration) *AdminPatchOAuth2ClientParams {
	return &AdminPatchOAuth2ClientParams{
		timeout: timeout,
	}
}

// NewAdminPatchOAuth2ClientParamsWithContext creates a new AdminPatchOAuth2ClientParams object
// with the ability to set a context for a request.
func NewAdminPatchOAuth2ClientParamsWithContext(ctx context.Context) *AdminPatchOAuth2ClientParams {
	return &AdminPatchOAuth2ClientParams{
		Context: ctx,
	}
}

// NewAdminPatchOAuth2ClientParamsWithHTTPClient creates a new AdminPatchOAuth2ClientParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminPatchOAuth2ClientParamsWithHTTPClient(client *http.Client) *AdminPatchOAuth2ClientParams {
	return &AdminPatchOAuth2ClientParams{
		HTTPClient: client,
	}
}

/* AdminPatchOAuth2ClientParams contains all the parameters to send to the API endpoint
   for the admin patch o auth2 client operation.

   Typically these are written to a http.Request.
*/
type AdminPatchOAuth2ClientParams struct {

	// Body.
	Body models.JSONPatchDocument

	/* ID.

	   The id of the OAuth 2.0 Client.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin patch o auth2 client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminPatchOAuth2ClientParams) WithDefaults() *AdminPatchOAuth2ClientParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin patch o auth2 client params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminPatchOAuth2ClientParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) WithTimeout(timeout time.Duration) *AdminPatchOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) WithContext(ctx context.Context) *AdminPatchOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) WithHTTPClient(client *http.Client) *AdminPatchOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) WithBody(body models.JSONPatchDocument) *AdminPatchOAuth2ClientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) SetBody(body models.JSONPatchDocument) {
	o.Body = body
}

// WithID adds the id to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) WithID(id string) *AdminPatchOAuth2ClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the admin patch o auth2 client params
func (o *AdminPatchOAuth2ClientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AdminPatchOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
