// Code generated by go-swagger; DO NOT EDIT.

package employment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/models"
)

// GetEmploymentReader is a Reader for the GetEmployment structure.
type GetEmploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetEmploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetEmploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetEmploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEmploymentOK creates a GetEmploymentOK with default headers values
func NewGetEmploymentOK() *GetEmploymentOK {
	return &GetEmploymentOK{}
}

/* GetEmploymentOK describes a response with status code 200, with default header values.

A employment
*/
type GetEmploymentOK struct {
	Payload *models.EmploymentDto
}

func (o *GetEmploymentOK) Error() string {
	return fmt.Sprintf("[GET /v1/employment/{id}][%d] getEmploymentOK  %+v", 200, o.Payload)
}
func (o *GetEmploymentOK) GetPayload() *models.EmploymentDto {
	return o.Payload
}

func (o *GetEmploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmploymentDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmploymentForbidden creates a GetEmploymentForbidden with default headers values
func NewGetEmploymentForbidden() *GetEmploymentForbidden {
	return &GetEmploymentForbidden{}
}

/* GetEmploymentForbidden describes a response with status code 403, with default header values.

Generic error message returned as a string
*/
type GetEmploymentForbidden struct {
	Payload models.GenericError
}

func (o *GetEmploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/employment/{id}][%d] getEmploymentForbidden  %+v", 403, o.Payload)
}
func (o *GetEmploymentForbidden) GetPayload() models.GenericError {
	return o.Payload
}

func (o *GetEmploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmploymentNotFound creates a GetEmploymentNotFound with default headers values
func NewGetEmploymentNotFound() *GetEmploymentNotFound {
	return &GetEmploymentNotFound{}
}

/* GetEmploymentNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type GetEmploymentNotFound struct {
	Payload models.GenericError
}

func (o *GetEmploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/employment/{id}][%d] getEmploymentNotFound  %+v", 404, o.Payload)
}
func (o *GetEmploymentNotFound) GetPayload() models.GenericError {
	return o.Payload
}

func (o *GetEmploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmploymentInternalServerError creates a GetEmploymentInternalServerError with default headers values
func NewGetEmploymentInternalServerError() *GetEmploymentInternalServerError {
	return &GetEmploymentInternalServerError{}
}

/* GetEmploymentInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type GetEmploymentInternalServerError struct {
	Payload models.GenericError
}

func (o *GetEmploymentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/employment/{id}][%d] getEmploymentInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEmploymentInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *GetEmploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
