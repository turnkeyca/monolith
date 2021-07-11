package employment

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
	Id string `json:"id"`
}

// swagger:parameters getEmploymentByUserId
//lint:ignore U1000 for docs
type employmentUserIdParamsWrapper struct {
	// The user id
	// in: query
	// required: true
	UserId string `json:"userId"`
}
