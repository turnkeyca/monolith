// Package classification of Turnkey API
//
// Documentation for Turnkey API
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package employment

import "github.com/google/uuid"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response employmentErrorResponse
//lint:ignore U1000 for docs
type employmentErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response employmentErrorValidation
//lint:ignore U1000 for docs
type employmentErrorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
}

// A employment
// swagger:response employmentResponse
//lint:ignore U1000 for docs
type employmentResponseWrapper struct {
	// A employment
	// in: body
	Body EmploymentDto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters updateEmployment getEmployment deleteEmployment
//lint:ignore U1000 for docs
type employmentIdParamsWrapper struct {
	// The id of the employment for which the operation relates
	// in: path
	// required: true
	Id uuid.UUID `json:"id"`
}
