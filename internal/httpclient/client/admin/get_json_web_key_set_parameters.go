// Code generated by go-swagger; DO NOT EDIT.

package admin

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
)

// NewGetJSONWebKeySetParams creates a new GetJSONWebKeySetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetJSONWebKeySetParams() *GetJSONWebKeySetParams {
	return &GetJSONWebKeySetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetJSONWebKeySetParamsWithTimeout creates a new GetJSONWebKeySetParams object
// with the ability to set a timeout on a request.
func NewGetJSONWebKeySetParamsWithTimeout(timeout time.Duration) *GetJSONWebKeySetParams {
	return &GetJSONWebKeySetParams{
		timeout: timeout,
	}
}

// NewGetJSONWebKeySetParamsWithContext creates a new GetJSONWebKeySetParams object
// with the ability to set a context for a request.
func NewGetJSONWebKeySetParamsWithContext(ctx context.Context) *GetJSONWebKeySetParams {
	return &GetJSONWebKeySetParams{
		Context: ctx,
	}
}

// NewGetJSONWebKeySetParamsWithHTTPClient creates a new GetJSONWebKeySetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetJSONWebKeySetParamsWithHTTPClient(client *http.Client) *GetJSONWebKeySetParams {
	return &GetJSONWebKeySetParams{
		HTTPClient: client,
	}
}

/* GetJSONWebKeySetParams contains all the parameters to send to the API endpoint
   for the get Json web key set operation.

   Typically these are written to a http.Request.
*/
type GetJSONWebKeySetParams struct {

	/* Set.

	   The set
	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get Json web key set params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJSONWebKeySetParams) WithDefaults() *GetJSONWebKeySetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get Json web key set params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetJSONWebKeySetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get Json web key set params
func (o *GetJSONWebKeySetParams) WithTimeout(timeout time.Duration) *GetJSONWebKeySetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get Json web key set params
func (o *GetJSONWebKeySetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get Json web key set params
func (o *GetJSONWebKeySetParams) WithContext(ctx context.Context) *GetJSONWebKeySetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get Json web key set params
func (o *GetJSONWebKeySetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get Json web key set params
func (o *GetJSONWebKeySetParams) WithHTTPClient(client *http.Client) *GetJSONWebKeySetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get Json web key set params
func (o *GetJSONWebKeySetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSet adds the set to the get Json web key set params
func (o *GetJSONWebKeySetParams) WithSet(set string) *GetJSONWebKeySetParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the get Json web key set params
func (o *GetJSONWebKeySetParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *GetJSONWebKeySetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param set
	if err := r.SetPathParam("set", o.Set); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
