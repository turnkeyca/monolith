// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/models"
)

// DeletePetReader is a Reader for the DeletePet structure.
type DeletePetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeletePetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeletePetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePetNoContent creates a DeletePetNoContent with default headers values
func NewDeletePetNoContent() *DeletePetNoContent {
	return &DeletePetNoContent{}
}

/* DeletePetNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type DeletePetNoContent struct {
}

func (o *DeletePetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/pet/{id}][%d] deletePetNoContent ", 204)
}

func (o *DeletePetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePetForbidden creates a DeletePetForbidden with default headers values
func NewDeletePetForbidden() *DeletePetForbidden {
	return &DeletePetForbidden{}
}

/* DeletePetForbidden describes a response with status code 403, with default header values.

Generic error message returned as a string
*/
type DeletePetForbidden struct {
	Payload models.GenericError
}

func (o *DeletePetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /v1/pet/{id}][%d] deletePetForbidden  %+v", 403, o.Payload)
}
func (o *DeletePetForbidden) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeletePetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePetNotFound creates a DeletePetNotFound with default headers values
func NewDeletePetNotFound() *DeletePetNotFound {
	return &DeletePetNotFound{}
}

/* DeletePetNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type DeletePetNotFound struct {
	Payload models.GenericError
}

func (o *DeletePetNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/pet/{id}][%d] deletePetNotFound  %+v", 404, o.Payload)
}
func (o *DeletePetNotFound) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeletePetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePetInternalServerError creates a DeletePetInternalServerError with default headers values
func NewDeletePetInternalServerError() *DeletePetInternalServerError {
	return &DeletePetInternalServerError{}
}

/* DeletePetInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type DeletePetInternalServerError struct {
	Payload models.GenericError
}

func (o *DeletePetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/pet/{id}][%d] deletePetInternalServerError  %+v", 500, o.Payload)
}
func (o *DeletePetInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeletePetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
