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
package reference

import "github.com/google/uuid"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response referenceErrorResponse
//lint:ignore U1000 for docs
type referenceErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response referenceErrorValidation
//lint:ignore U1000 for docs
type referenceErrorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
}

// A reference
// swagger:response referenceResponse
//lint:ignore U1000 for docs
type referenceResponseWrapper struct {
	// A reference
	// in: body
	Body Dto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters updateReference getReference deleteReference
//lint:ignore U1000 for docs
type referenceIdParamsWrapper struct {
	// The id of the reference for which the operation relates
	// in: path
	// required: true
	Id uuid.UUID `json:"id"`
}
