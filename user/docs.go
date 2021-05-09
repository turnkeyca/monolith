// Package classification of Turnkey API
//
// Documentation for Turnkey API
//
//	Schemes: http, https
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
package user

import "github.com/google/uuid"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response errorResponse
//lint:ignore U1000 for docs
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
//lint:ignore U1000 for docs
type errorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
}

// A user
// swagger:response userResponse
//lint:ignore U1000 for docs
type userResponseWrapper struct {
	// All current products
	// in: body
	Body Dto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters updateUser getUser deleteUser
//lint:ignore U1000 for docs
type userIdParamsWrapper struct {
	// The id of the user for which the operation relates
	// in: path
	// required: true
	Id uuid.UUID `json:"id"`
}
