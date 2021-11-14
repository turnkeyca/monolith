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

// GetEmploymentsByUserIDReader is a Reader for the GetEmploymentsByUserID structure.
type GetEmploymentsByUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmploymentsByUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEmploymentsByUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetEmploymentsByUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEmploymentsByUserIDOK creates a GetEmploymentsByUserIDOK with default headers values
func NewGetEmploymentsByUserIDOK() *GetEmploymentsByUserIDOK {
	return &GetEmploymentsByUserIDOK{}
}

/* GetEmploymentsByUserIDOK describes a response with status code 200, with default header values.

A list of employment
*/
type GetEmploymentsByUserIDOK struct {
	Payload []*models.EmploymentDto
}

func (o *GetEmploymentsByUserIDOK) Error() string {
	return fmt.Sprintf("[GET /v1/employment][%d] getEmploymentsByUserIdOK  %+v", 200, o.Payload)
}
func (o *GetEmploymentsByUserIDOK) GetPayload() []*models.EmploymentDto {
	return o.Payload
}

func (o *GetEmploymentsByUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmploymentsByUserIDInternalServerError creates a GetEmploymentsByUserIDInternalServerError with default headers values
func NewGetEmploymentsByUserIDInternalServerError() *GetEmploymentsByUserIDInternalServerError {
	return &GetEmploymentsByUserIDInternalServerError{}
}

/* GetEmploymentsByUserIDInternalServerError describes a response with status code 500, with default header values.

Generic error message returned as a string
*/
type GetEmploymentsByUserIDInternalServerError struct {
	Payload models.GenericError
}

func (o *GetEmploymentsByUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/employment][%d] getEmploymentsByUserIdInternalServerError  %+v", 500, o.Payload)
}
func (o *GetEmploymentsByUserIDInternalServerError) GetPayload() models.GenericError {
	return o.Payload
}

func (o *GetEmploymentsByUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}