package employment

import "github.com/turnkeyca/monolith/util"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response employmentErrorResponse
//lint:ignore U1000 for docs
type employmentErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body util.GenericError
}

// A employment
// swagger:response employmentResponse
//lint:ignore U1000 for docs
type employmentResponseWrapper struct {
	// A employment
	// in: body
	Body EmploymentDto
}

// A list of employment
// swagger:response employmentsResponse
//lint:ignore U1000 for docs
type employmentsResponseWrapper struct {
	// A list of employment
	// in: body
	Body []EmploymentDto
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
	Id string `json:"ID"`
}

// swagger:parameters getEmploymentsByUserId
//lint:ignore U1000 for docs
type employmentUserIdParamsWrapper struct {
	// The user id
	// in: query
	// required: true
	UserId string `json:"userID"`
}

// swagger:parameters updateEmployment createEmployment
//lint:ignore U1000 for docs
type employmentParamsWrapper struct {
	// in: body
	// required: true
	Body EmploymentDto
}

// swagger:parameters updateEmployment createEmployment getEmploymentsByUserId getEmployment deleteEmployment
//lint:ignore U1000 for docs
type authHeaderWrapper struct {
	// in: header
	// required: true
	Token string
}
