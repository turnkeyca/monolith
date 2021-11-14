// Code generated by go-swagger; DO NOT EDIT.

package auth

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/models"
)

// RegisterNewTokenReader is a Reader for the RegisterNewToken structure.
type RegisterNewTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterNewTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterNewTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewRegisterNewTokenInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterNewTokenOK creates a RegisterNewTokenOK with default headers values
func NewRegisterNewTokenOK() *RegisterNewTokenOK {
	return &RegisterNewTokenOK{}
}

/* RegisterNewTokenOK describes a response with status code 200, with default header values.

new user id
*/
type RegisterNewTokenOK struct {
	Payload *models.TokenDto
}

func (o *RegisterNewTokenOK) Error() string {
	return fmt.Sprintf("[POST /v1/auth/registertoken][%d] registerNewTokenOK  %+v", 200, o.Payload)
}
func (o *RegisterNewTokenOK) GetPayload() *models.TokenDto {
	return o.Payload
}

func (o *RegisterNewTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TokenDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterNewTokenInternalServerError creates a RegisterNewTokenInternalServerError with default headers values
func NewRegisterNewTokenInternalServerError() *RegisterNewTokenInternalServerError {
	return &RegisterNewTokenInternalServerError{}
}

/* RegisterNewTokenInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type RegisterNewTokenInternalServerError struct {
	Payload models.GenericError
}

func (o *RegisterNewTokenInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/auth/registertoken][%d] registerNewTokenInternalServerError  %+v", 500, o.Payload)
}
func (o *RegisterNewTokenInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *RegisterNewTokenInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
