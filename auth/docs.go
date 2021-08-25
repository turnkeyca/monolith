package auth

// Generic error message returned as a string
// swagger:response authErrorResponse
//lint:ignore U1000 for docs
type authErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// new user id
// swagger:response userIdResponse
//lint:ignore U1000 for docs
type userIdResponseWrapper struct {
	// user id
	// in: body
	Body UserId
}

// swagger:parameters registerToken
//lint:ignore U1000 for docs
type registerTokenParamsWrapper struct {
	// in: body
	// required: true
	Body RegisterTokenDto
}
