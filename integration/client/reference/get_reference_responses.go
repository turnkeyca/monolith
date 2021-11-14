// Code generated by go-swagger; DO NOT EDIT.

package reference

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/models"
)

// GetReferenceReader is a Reader for the GetReference structure.
type GetReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetReferenceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetReferenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetReferenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetReferenceOK creates a GetReferenceOK with default headers values
func NewGetReferenceOK() *GetReferenceOK {
	return &GetReferenceOK{}
}

/* GetReferenceOK describes a response with status code 200, with default header values.

A reference
*/
type GetReferenceOK struct {
	Payload *models.ReferenceDto
}

func (o *GetReferenceOK) Error() string {
	return fmt.Sprintf("[GET /v1/reference/{id}][%d] getReferenceOK  %+v", 200, o.Payload)
}
func (o *GetReferenceOK) GetPayload() *models.ReferenceDto {
	return o.Payload
}

func (o *GetReferenceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReferenceDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReferenceNotFound creates a GetReferenceNotFound with default headers values
func NewGetReferenceNotFound() *GetReferenceNotFound {
	return &GetReferenceNotFound{}
}

/* GetReferenceNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type GetReferenceNotFound struct {
	Payload models.GenericError
}

func (o *GetReferenceNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/reference/{id}][%d] getReferenceNotFound  %+v", 404, o.Payload)
}
func (o *GetReferenceNotFound) GetPayload() models.GenericError {
	return o.Payload
}

func (o *GetReferenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReferenceInternalServerError creates a GetReferenceInternalServerError with default headers values
func NewGetReferenceInternalServerError() *GetReferenceInternalServerError {
	return &GetReferenceInternalServerError{}
}

/* GetReferenceInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type GetReferenceInternalServerError struct {
	Payload models.GenericError
}

func (o *GetReferenceInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/reference/{id}][%d] getReferenceInternalServerError  %+v", 500, o.Payload)
}
func (o *GetReferenceInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *GetReferenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}