package authenticator

import "github.com/turnkeyca/monolith/util"

// Generic error message returned as a string
// swagger:response authenticationErrorResponse
//lint:ignore U1000 for docs
type authErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body util.GenericError
}

// new user id
// swagger:response tokenResponse
//lint:ignore U1000 for docs
type tokenResponseWrapper struct {
	// user id
	// in: body
	Body TokenDto
}

// swagger:parameters registerNewToken
//lint:ignore U1000 for docs
type registerTokenParamsWrapper struct {
	// in: body
	// required: true
	Body RegisterTokenDto
}
