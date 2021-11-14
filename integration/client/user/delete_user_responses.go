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

// DeleteUserReader is a Reader for the DeleteUser structure.
type DeleteUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteUserNoContent creates a DeleteUserNoContent with default headers values
func NewDeleteUserNoContent() *DeleteUserNoContent {
	return &DeleteUserNoContent{}
}

/* DeleteUserNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type DeleteUserNoContent struct {
}

func (o *DeleteUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/user/{id}][%d] deleteUserNoContent ", 204)
}

func (o *DeleteUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserInternalServerError creates a DeleteUserInternalServerError with default headers values
func NewDeleteUserInternalServerError() *DeleteUserInternalServerError {
	return &DeleteUserInternalServerError{}
}

/* DeleteUserInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type DeleteUserInternalServerError struct {
	Payload models.GenericError
}

func (o *DeleteUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/user/{id}][%d] deleteUserInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteUserInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}