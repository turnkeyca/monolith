// Code generated by go-swagger; DO NOT EDIT.

package user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/models"
)

// UpdateUserReader is a Reader for the UpdateUser structure.
type UpdateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewUpdateUserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateUserNoContent creates a UpdateUserNoContent with default headers values
func NewUpdateUserNoContent() *UpdateUserNoContent {
	return &UpdateUserNoContent{}
}

/* UpdateUserNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type UpdateUserNoContent struct {
}

func (o *UpdateUserNoContent) Error() string {
	return fmt.Sprintf("[PUT /v1/user/{id}][%d] updateUserNoContent ", 204)
}

func (o *UpdateUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUserBadRequest creates a UpdateUserBadRequest with default headers values
func NewUpdateUserBadRequest() *UpdateUserBadRequest {
	return &UpdateUserBadRequest{}
}

/* UpdateUserBadRequest describes a response with status code 400, with default header values.

Generic error message returned as a string
*/
type UpdateUserBadRequest struct {
	Payload models.GenericError
}

func (o *UpdateUserBadRequest) Error() string {
	return fmt.Sprintf("[PUT /v1/user/{id}][%d] updateUserBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateUserBadRequest) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserForbidden creates a UpdateUserForbidden with default headers values
func NewUpdateUserForbidden() *UpdateUserForbidden {
	return &UpdateUserForbidden{}
}

/* UpdateUserForbidden describes a response with status code 403, with default header values.

Generic error message returned as a string
*/
type UpdateUserForbidden struct {
	Payload models.GenericError
}

func (o *UpdateUserForbidden) Error() string {
	return fmt.Sprintf("[PUT /v1/user/{id}][%d] updateUserForbidden  %+v", 403, o.Payload)
}
func (o *UpdateUserForbidden) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserUnprocessableEntity creates a UpdateUserUnprocessableEntity with default headers values
func NewUpdateUserUnprocessableEntity() *UpdateUserUnprocessableEntity {
	return &UpdateUserUnprocessableEntity{}
}

/* UpdateUserUnprocessableEntity describes a response with status code 422, with default header values.

Generic error message returned as a string
*/
type UpdateUserUnprocessableEntity struct {
	Payload models.GenericError
}

func (o *UpdateUserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/user/{id}][%d] updateUserUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *UpdateUserUnprocessableEntity) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateUserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserInternalServerError creates a UpdateUserInternalServerError with default headers values
func NewUpdateUserInternalServerError() *UpdateUserInternalServerError {
	return &UpdateUserInternalServerError{}
}

/* UpdateUserInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type UpdateUserInternalServerError struct {
	Payload models.GenericError
}

func (o *UpdateUserInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/user/{id}][%d] updateUserInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateUserInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *UpdateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
