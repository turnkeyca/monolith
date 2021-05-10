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
package shorturl

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handers

// Generic error message returned as a string
// swagger:response shortUrlErrorResponse
//lint:ignore U1000 for docs
type shortUrlErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// A short url
// swagger:response shortUrlResponse
//lint:ignore U1000 for docs
type shortUrlResponseWrapper struct {
	// A short url
	// in: body
	Body Dto
}
