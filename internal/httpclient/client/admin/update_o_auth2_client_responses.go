// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/hydra/internal/httpclient/models"
)

// UpdateOAuth2ClientReader is a Reader for the UpdateOAuth2Client structure.
type UpdateOAuth2ClientReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOAuth2ClientReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOAuth2ClientOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateOAuth2ClientDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateOAuth2ClientOK creates a UpdateOAuth2ClientOK with default headers values
func NewUpdateOAuth2ClientOK() *UpdateOAuth2ClientOK {
	return &UpdateOAuth2ClientOK{}
}

/* UpdateOAuth2ClientOK describes a response with status code 200, with default header values.

oAuth2Client
*/
type UpdateOAuth2ClientOK struct {
	Payload *models.OAuth2Client
}

func (o *UpdateOAuth2ClientOK) Error() string {
	return fmt.Sprintf("[PUT /clients/{id}][%d] updateOAuth2ClientOK  %+v", 200, o.Payload)
}
func (o *UpdateOAuth2ClientOK) GetPayload() *models.OAuth2Client {
	return o.Payload
}

func (o *UpdateOAuth2ClientOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OAuth2Client)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateOAuth2ClientDefault creates a UpdateOAuth2ClientDefault with default headers values
func NewUpdateOAuth2ClientDefault(code int) *UpdateOAuth2ClientDefault {
	return &UpdateOAuth2ClientDefault{
		_statusCode: code,
	}
}

/* UpdateOAuth2ClientDefault describes a response with status code -1, with default header values.

jsonError
*/
type UpdateOAuth2ClientDefault struct {
	_statusCode int

	Payload *models.JSONError
}

// Code gets the status code for the update o auth2 client default response
func (o *UpdateOAuth2ClientDefault) Code() int {
	return o._statusCode
}

func (o *UpdateOAuth2ClientDefault) Error() string {
	return fmt.Sprintf("[PUT /clients/{id}][%d] updateOAuth2Client default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateOAuth2ClientDefault) GetPayload() *models.JSONError {
	return o.Payload
}

func (o *UpdateOAuth2ClientDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
