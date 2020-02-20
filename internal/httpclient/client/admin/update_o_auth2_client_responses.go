// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/hydra/internal/httpclient/models"
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
	case 500:
		result := NewUpdateOAuth2ClientInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateOAuth2ClientOK creates a UpdateOAuth2ClientOK with default headers values
func NewUpdateOAuth2ClientOK() *UpdateOAuth2ClientOK {
	return &UpdateOAuth2ClientOK{}
}

/*UpdateOAuth2ClientOK handles this case with default header values.

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

// NewUpdateOAuth2ClientInternalServerError creates a UpdateOAuth2ClientInternalServerError with default headers values
func NewUpdateOAuth2ClientInternalServerError() *UpdateOAuth2ClientInternalServerError {
	return &UpdateOAuth2ClientInternalServerError{}
}

/*UpdateOAuth2ClientInternalServerError handles this case with default header values.

genericError
*/
type UpdateOAuth2ClientInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateOAuth2ClientInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /clients/{id}][%d] updateOAuth2ClientInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateOAuth2ClientInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateOAuth2ClientInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
