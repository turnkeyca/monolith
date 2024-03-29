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

// UpdateEmploymentReader is a Reader for the UpdateEmployment structure.
type UpdateEmploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateEmploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateEmploymentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateEmploymentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateEmploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateEmploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateEmploymentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateEmploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateEmploymentNoContent creates a UpdateEmploymentNoContent with default headers values
func NewUpdateEmploymentNoContent() *UpdateEmploymentNoContent {
	return &UpdateEmploymentNoContent{}
}

/* UpdateEmploymentNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type UpdateEmploymentNoContent struct {
}

func (o *UpdateEmploymentNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/employment/{id}][%d] updateEmploymentNoContent ", 204)
}

func (o *UpdateEmploymentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateEmploymentBadRequest creates a UpdateEmploymentBadRequest with default headers values
func NewUpdateEmploymentBadRequest() *UpdateEmploymentBadRequest {
	return &UpdateEmploymentBadRequest{}
}

/* UpdateEmploymentBadRequest describes a response with status code 400, with default header values.

Generic error message returned as a string
*/
type UpdateEmploymentBadRequest struct {
	Payload models.GenericError
}

func (o *UpdateEmploymentBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/employment/{id}][%d] updateEmploymentBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateEmploymentBadRequest) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateEmploymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEmploymentForbidden creates a UpdateEmploymentForbidden with default headers values
func NewUpdateEmploymentForbidden() *UpdateEmploymentForbidden {
	return &UpdateEmploymentForbidden{}
}

/* UpdateEmploymentForbidden describes a response with status code 403, with default header values.

Generic error message returned as a string
*/
type UpdateEmploymentForbidden struct {
	Payload models.GenericError
}

func (o *UpdateEmploymentForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/employment/{id}][%d] updateEmploymentForbidden  %+v", 403, o.Payload)
}
func (o *UpdateEmploymentForbidden) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateEmploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEmploymentNotFound creates a UpdateEmploymentNotFound with default headers values
func NewUpdateEmploymentNotFound() *UpdateEmploymentNotFound {
	return &UpdateEmploymentNotFound{}
}

/* UpdateEmploymentNotFound describes a response with status code 404, with default header values.

Generic error message returned as a string
*/
type UpdateEmploymentNotFound struct {
	Payload models.GenericError
}

func (o *UpdateEmploymentNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/employment/{id}][%d] updateEmploymentNotFound  %+v", 404, o.Payload)
}
func (o *UpdateEmploymentNotFound) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateEmploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEmploymentUnprocessableEntity creates a UpdateEmploymentUnprocessableEntity with default headers values
func NewUpdateEmploymentUnprocessableEntity() *UpdateEmploymentUnprocessableEntity {
	return &UpdateEmploymentUnprocessableEntity{}
}

/* UpdateEmploymentUnprocessableEntity describes a response with status code 422, with default header values.

Generic error message returned as a string
*/
type UpdateEmploymentUnprocessableEntity struct {
	Payload models.GenericError
}

func (o *UpdateEmploymentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/employment/{id}][%d] updateEmploymentUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateEmploymentUnprocessableEntity) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateEmploymentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateEmploymentInternalServerError creates a UpdateEmploymentInternalServerError with default headers values
func NewUpdateEmploymentInternalServerError() *UpdateEmploymentInternalServerError {
	return &UpdateEmploymentInternalServerError{}
}

/* UpdateEmploymentInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type UpdateEmploymentInternalServerError struct {
	Payload models.GenericError
}

func (o *UpdateEmploymentInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/employment/{id}][%d] updateEmploymentInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateEmploymentInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateEmploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
