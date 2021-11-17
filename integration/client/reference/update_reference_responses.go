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

// UpdateReferenceReader is a Reader for the UpdateReference structure.
type UpdateReferenceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReferenceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateReferenceNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateReferenceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateReferenceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateReferenceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateReferenceUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateReferenceInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateReferenceNoContent creates a UpdateReferenceNoContent with default headers values
func NewUpdateReferenceNoContent() *UpdateReferenceNoContent {
	return &UpdateReferenceNoContent{}
}

/* UpdateReferenceNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type UpdateReferenceNoContent struct {
}

func (o *UpdateReferenceNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/reference/{id}][%d] updateReferenceNoContent ", 204)
}

func (o *UpdateReferenceNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateReferenceBadRequest creates a UpdateReferenceBadRequest with default headers values
func NewUpdateReferenceBadRequest() *UpdateReferenceBadRequest {
	return &UpdateReferenceBadRequest{}
}

/* UpdateReferenceBadRequest describes a response with status code 400, with default header values.

Generic error message returned as a string
*/
type UpdateReferenceBadRequest struct {
	Payload models.GenericError
}

func (o *UpdateReferenceBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/reference/{id}][%d] updateReferenceBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateReferenceBadRequest) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateReferenceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReferenceForbidden creates a UpdateReferenceForbidden with default headers values
func NewUpdateReferenceForbidden() *UpdateReferenceForbidden {
	return &UpdateReferenceForbidden{}
}

/* UpdateReferenceForbidden describes a response with status code 403, with default header values.

Generic error message returned as a string
*/
type UpdateReferenceForbidden struct {
	Payload models.GenericError
}

func (o *UpdateReferenceForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/reference/{id}][%d] updateReferenceForbidden  %+v", 403, o.Payload)
}
func (o *UpdateReferenceForbidden) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateReferenceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReferenceNotFound creates a UpdateReferenceNotFound with default headers values
func NewUpdateReferenceNotFound() *UpdateReferenceNotFound {
	return &UpdateReferenceNotFound{}
}

/* UpdateReferenceNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type UpdateReferenceNotFound struct {
	Payload models.GenericError
}

func (o *UpdateReferenceNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/reference/{id}][%d] updateReferenceNotFound  %+v", 404, o.Payload)
}
func (o *UpdateReferenceNotFound) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateReferenceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReferenceUnprocessableEntity creates a UpdateReferenceUnprocessableEntity with default headers values
func NewUpdateReferenceUnprocessableEntity() *UpdateReferenceUnprocessableEntity {
	return &UpdateReferenceUnprocessableEntity{}
}

/* UpdateReferenceUnprocessableEntity describes a response with status code 422, with default header values.

Generic error message returned as a string
*/
type UpdateReferenceUnprocessableEntity struct {
	Payload models.GenericError
}

func (o *UpdateReferenceUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/reference/{id}][%d] updateReferenceUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateReferenceUnprocessableEntity) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateReferenceUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReferenceInternalServerError creates a UpdateReferenceInternalServerError with default headers values
func NewUpdateReferenceInternalServerError() *UpdateReferenceInternalServerError {
	return &UpdateReferenceInternalServerError{}
}

/* UpdateReferenceInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type UpdateReferenceInternalServerError struct {
	Payload models.GenericError
}

func (o *UpdateReferenceInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/reference/{id}][%d] updateReferenceInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateReferenceInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateReferenceInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
