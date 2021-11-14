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

// DeleteEmploymentReader is a Reader for the DeleteEmployment structure.
type DeleteEmploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEmploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteEmploymentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteEmploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteEmploymentNoContent creates a DeleteEmploymentNoContent with default headers values
func NewDeleteEmploymentNoContent() *DeleteEmploymentNoContent {
	return &DeleteEmploymentNoContent{}
}

/* DeleteEmploymentNoContent describes a response with status code 204, with default header values.

No content is returned by this API endpoint
*/
type DeleteEmploymentNoContent struct {
}

func (o *DeleteEmploymentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /v1/employment/{id}][%d] deleteEmploymentNoContent ", 204)
}

func (o *DeleteEmploymentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmploymentInternalServerError creates a DeleteEmploymentInternalServerError with default headers values
func NewDeleteEmploymentInternalServerError() *DeleteEmploymentInternalServerError {
	return &DeleteEmploymentInternalServerError{}
}

/* DeleteEmploymentInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type DeleteEmploymentInternalServerError struct {
	Payload models.GenericError
}

func (o *DeleteEmploymentInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /v1/employment/{id}][%d] deleteEmploymentInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteEmploymentInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *DeleteEmploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
