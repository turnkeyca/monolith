// Code generated by go-swagger; DO NOT EDIT.

package permission

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/turnkeyca/monolith/integration/models"
)

// DeletePermissionReader is a Reader for the DeletePermission structure.
type DeletePermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeletePermissionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewDeletePermissionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeletePermissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeletePermissionNoContent creates a DeletePermissionNoContent with default headers values
func NewDeletePermissionNoContent() *DeletePermissionNoContent {
	return &DeletePermissionNoContent{}
}

/* DeletePermissionNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type DeletePermissionNoContent struct {
}

func (o *DeletePermissionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/permission/{id}][%d] deletePermissionNoContent ", 204)
}

func (o *DeletePermissionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePermissionConflict creates a DeletePermissionConflict with default headers values
func NewDeletePermissionConflict() *DeletePermissionConflict {
	return &DeletePermissionConflict{}
}

/* DeletePermissionConflict describes a response with status code 409, with default header values.

Generic error message returned as a string
*/
type DeletePermissionConflict struct {
	Payload models.GenericError
}

func (o *DeletePermissionConflict) Error() string {
	return fmt.Sprintf("[DELETE /v1/permission/{id}][%d] deletePermissionConflict  %+v", 409, o.Payload)
}
func (o *DeletePermissionConflict) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeletePermissionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeletePermissionInternalServerError creates a DeletePermissionInternalServerError with default headers values
func NewDeletePermissionInternalServerError() *DeletePermissionInternalServerError {
	return &DeletePermissionInternalServerError{}
}

/* DeletePermissionInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type DeletePermissionInternalServerError struct {
	Payload models.GenericError
}

func (o *DeletePermissionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/permission/{id}][%d] deletePermissionInternalServerError  %+v", 500, o.Payload)
}
func (o *DeletePermissionInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeletePermissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}