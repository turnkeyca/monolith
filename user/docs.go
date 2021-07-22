package user

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response userErrorResponse
//lint:ignore U1000 for docs
type userErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response userErrorValidation
//lint:ignore U1000 for docs
type userErrorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
}

// A user
// swagger:response userResponse
//lint:ignore U1000 for docs
type userResponseWrapper struct {
	// A user
	// in: body
	Body UserDto
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
	Id string `json:"id"`
}

// swagger:parameters updateUser createUser
//lint:ignore U1000 for docs
type userParamsWrapper struct {
	// in: body
	// required: true
	Body UserDto
}
