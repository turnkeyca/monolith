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

// CreatePetReader is a Reader for the CreatePet structure.
type CreatePetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCreatePetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewCreatePetUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreatePetNoContent creates a CreatePetNoContent with default headers values
func NewCreatePetNoContent() *CreatePetNoContent {
	return &CreatePetNoContent{}
}

/* CreatePetNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type CreatePetNoContent struct {
}

func (o *CreatePetNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/pet][%d] createPetNoContent ", 204)
}

func (o *CreatePetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePetBadRequest creates a CreatePetBadRequest with default headers values
func NewCreatePetBadRequest() *CreatePetBadRequest {
	return &CreatePetBadRequest{}
}

/* CreatePetBadRequest describes a response with status code 400, with default header values.

Generic error message returned as a string
*/
type CreatePetBadRequest struct {
	Payload models.GenericError
}

func (o *CreatePetBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/pet][%d] createPetBadRequest  %+v", 400, o.Payload)
}
func (o *CreatePetBadRequest) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreatePetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePetForbidden creates a CreatePetForbidden with default headers values
func NewCreatePetForbidden() *CreatePetForbidden {
	return &CreatePetForbidden{}
}

/* CreatePetForbidden describes a response with status code 403, with default header values.

Generic error message returned as a string
*/
type CreatePetForbidden struct {
	Payload models.GenericError
}

func (o *CreatePetForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/pet][%d] createPetForbidden  %+v", 403, o.Payload)
}
func (o *CreatePetForbidden) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreatePetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePetUnprocessableEntity creates a CreatePetUnprocessableEntity with default headers values
func NewCreatePetUnprocessableEntity() *CreatePetUnprocessableEntity {
	return &CreatePetUnprocessableEntity{}
}

/* CreatePetUnprocessableEntity describes a response with status code 422, with default header values.

Generic error message returned as a string
*/
type CreatePetUnprocessableEntity struct {
	Payload models.GenericError
}

func (o *CreatePetUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/pet][%d] createPetUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreatePetUnprocessableEntity) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreatePetUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePetInternalServerError creates a CreatePetInternalServerError with default headers values
func NewCreatePetInternalServerError() *CreatePetInternalServerError {
	return &CreatePetInternalServerError{}
}

/* CreatePetInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type CreatePetInternalServerError struct {
	Payload models.GenericError
}

func (o *CreatePetInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/pet][%d] createPetInternalServerError  %+v", 500, o.Payload)
}
func (o *CreatePetInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *CreatePetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
