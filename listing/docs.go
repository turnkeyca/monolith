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
package listing

import "github.com/google/uuid"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response listingErrorResponse
//lint:ignore U1000 for docs
type listingErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response listingErrorValidation
//lint:ignore U1000 for docs
type listingErrorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
}

// A listing
// swagger:response listingResponse
//lint:ignore U1000 for docs
type listingResponseWrapper struct {
	// A listing
	// in: body
	Body Dto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters updateListing getListing deleteListing
//lint:ignore U1000 for docs
type listingIdParamsWrapper struct {
	// The id of the listing for which the operation relates
	// in: path
	// required: true
	Id uuid.UUID `json:"id"`
}
