// Code generated by go-swagger; DO NOT EDIT.

package public

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
)

// NewUserinfoParams creates a new UserinfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUserinfoParams() *UserinfoParams {
	return &UserinfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUserinfoParamsWithTimeout creates a new UserinfoParams object
// with the ability to set a timeout on a request.
func NewUserinfoParamsWithTimeout(timeout time.Duration) *UserinfoParams {
	return &UserinfoParams{
		timeout: timeout,
	}
}

// NewUserinfoParamsWithContext creates a new UserinfoParams object
// with the ability to set a context for a request.
func NewUserinfoParamsWithContext(ctx context.Context) *UserinfoParams {
	return &UserinfoParams{
		Context: ctx,
	}
}

// NewUserinfoParamsWithHTTPClient creates a new UserinfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewUserinfoParamsWithHTTPClient(client *http.Client) *UserinfoParams {
	return &UserinfoParams{
		HTTPClient: client,
	}
}

/* UserinfoParams contains all the parameters to send to the API endpoint
   for the userinfo operation.

   Typically these are written to a http.Request.
*/
type UserinfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the userinfo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserinfoParams) WithDefaults() *UserinfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the userinfo params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UserinfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the userinfo params
func (o *UserinfoParams) WithTimeout(timeout time.Duration) *UserinfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the userinfo params
func (o *UserinfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the userinfo params
func (o *UserinfoParams) WithContext(ctx context.Context) *UserinfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the userinfo params
func (o *UserinfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the userinfo params
func (o *UserinfoParams) WithHTTPClient(client *http.Client) *UserinfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the userinfo params
func (o *UserinfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *UserinfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
