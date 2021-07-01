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
package roommate

import "github.com/google/uuid"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response roommateErrorResponse
//lint:ignore U1000 for docs
type roommateErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response roommateErrorValidation
//lint:ignore U1000 for docs
type roommateErrorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
}

// A roommate
// swagger:response roommateResponse
//lint:ignore U1000 for docs
type roommateResponseWrapper struct {
	// A roommate
	// in: body
	Body RoommateDto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters updateRoommate getRoommate deleteRoommate
//lint:ignore U1000 for docs
type roommateIdParamsWrapper struct {
	// The id of the roommate for which the operation relates
	// in: path
	// required: true
	Id uuid.UUID `json:"id"`
}
