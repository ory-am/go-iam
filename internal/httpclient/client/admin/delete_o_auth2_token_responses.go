// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/client/v1"
	"github.com/ory/hydra/internal/httpclient/models"
)

// DeleteOAuth2TokenReader is a Reader for the DeleteOAuth2Token structure.
type DeleteOAuth2TokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOAuth2TokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOAuth2TokenNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOAuth2TokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDeleteOAuth2TokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteOAuth2TokenNoContent creates a DeleteOAuth2TokenNoContent with default headers values
func NewDeleteOAuth2TokenNoContent() *DeleteOAuth2TokenNoContent {
	return &DeleteOAuth2TokenNoContent{}
}

/* DeleteOAuth2TokenNoContent describes a response with status code 204, with default header values.

 Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteOAuth2TokenNoContent struct {
}

func (o *DeleteOAuth2TokenNoContent) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/tokens][%d] deleteOAuth2TokenNoContent ", 204)
}

func (o *DeleteOAuth2TokenNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOAuth2TokenUnauthorized creates a DeleteOAuth2TokenUnauthorized with default headers values
func NewDeleteOAuth2TokenUnauthorized() *DeleteOAuth2TokenUnauthorized {
	return &DeleteOAuth2TokenUnauthorized{}
}

/* DeleteOAuth2TokenUnauthorized describes a response with status code 401, with default header values.

oAuth2ApiError
*/
type DeleteOAuth2TokenUnauthorized struct {
	Payload *models.OAuth2APIError
}

func (o *DeleteOAuth2TokenUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/tokens][%d] deleteOAuth2TokenUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteOAuth2TokenUnauthorized) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *DeleteOAuth2TokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteOAuth2TokenDefault creates a DeleteOAuth2TokenDefault with default headers values
func NewDeleteOAuth2TokenDefault(code int) *DeleteOAuth2TokenDefault {
	return &DeleteOAuth2TokenDefault{
		_statusCode: code,
	}
}

/* DeleteOAuth2TokenDefault describes a response with status code -1, with default header values.

oAuth2ApiError
*/
type DeleteOAuth2TokenDefault struct {
	_statusCode int

	Payload *models.OAuth2APIError
}

// Code gets the status code for the delete o auth2 token default response
func (o *DeleteOAuth2TokenDefault) Code() int {
	return o._statusCode
}

func (o *DeleteOAuth2TokenDefault) Error() string {
	return fmt.Sprintf("[DELETE /oauth2/tokens][%d] deleteOAuth2Token default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteOAuth2TokenDefault) GetPayload() *models.OAuth2APIError {
	return o.Payload
}

func (o *DeleteOAuth2TokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2APIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
