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
)

// NewAdminGetOAuth2ConsentRequestParams creates a new AdminGetOAuth2ConsentRequestParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAdminGetOAuth2ConsentRequestParams() *AdminGetOAuth2ConsentRequestParams {
	return &AdminGetOAuth2ConsentRequestParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAdminGetOAuth2ConsentRequestParamsWithTimeout creates a new AdminGetOAuth2ConsentRequestParams object
// with the ability to set a timeout on a request.
func NewAdminGetOAuth2ConsentRequestParamsWithTimeout(timeout time.Duration) *AdminGetOAuth2ConsentRequestParams {
	return &AdminGetOAuth2ConsentRequestParams{
		timeout: timeout,
	}
}

// NewAdminGetOAuth2ConsentRequestParamsWithContext creates a new AdminGetOAuth2ConsentRequestParams object
// with the ability to set a context for a request.
func NewAdminGetOAuth2ConsentRequestParamsWithContext(ctx context.Context) *AdminGetOAuth2ConsentRequestParams {
	return &AdminGetOAuth2ConsentRequestParams{
		Context: ctx,
	}
}

// NewAdminGetOAuth2ConsentRequestParamsWithHTTPClient creates a new AdminGetOAuth2ConsentRequestParams object
// with the ability to set a custom HTTPClient for a request.
func NewAdminGetOAuth2ConsentRequestParamsWithHTTPClient(client *http.Client) *AdminGetOAuth2ConsentRequestParams {
	return &AdminGetOAuth2ConsentRequestParams{
		HTTPClient: client,
	}
}

/* AdminGetOAuth2ConsentRequestParams contains all the parameters to send to the API endpoint
   for the admin get o auth2 consent request operation.

   Typically these are written to a http.Request.
*/
type AdminGetOAuth2ConsentRequestParams struct {

	// ConsentChallenge.
	ConsentChallenge string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the admin get o auth2 consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGetOAuth2ConsentRequestParams) WithDefaults() *AdminGetOAuth2ConsentRequestParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the admin get o auth2 consent request params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AdminGetOAuth2ConsentRequestParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) WithTimeout(timeout time.Duration) *AdminGetOAuth2ConsentRequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) WithContext(ctx context.Context) *AdminGetOAuth2ConsentRequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) WithHTTPClient(client *http.Client) *AdminGetOAuth2ConsentRequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConsentChallenge adds the consentChallenge to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) WithConsentChallenge(consentChallenge string) *AdminGetOAuth2ConsentRequestParams {
	o.SetConsentChallenge(consentChallenge)
	return o
}

// SetConsentChallenge adds the consentChallenge to the admin get o auth2 consent request params
func (o *AdminGetOAuth2ConsentRequestParams) SetConsentChallenge(consentChallenge string) {
	o.ConsentChallenge = consentChallenge
}

// WriteToRequest writes these params to a swagger request
func (o *AdminGetOAuth2ConsentRequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param consent_challenge
	qrConsentChallenge := o.ConsentChallenge
	qConsentChallenge := qrConsentChallenge
	if qConsentChallenge != "" {

		if err := r.SetQueryParam("consent_challenge", qConsentChallenge); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
