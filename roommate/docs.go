package roommate

import "github.com/turnkeyca/monolith/util"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response roommateErrorResponse
//lint:ignore U1000 for docs
type roommateErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body util.GenericError
}

// A roommate
// swagger:response roommateResponse
//lint:ignore U1000 for docs
type roommateResponseWrapper struct {
	// A roommate
	// in: body
	Body RoommateDto
}

// A list of roommates
// swagger:response roommatesResponse
//lint:ignore U1000 for docs
type roommatesResponseWrapper struct {
	// A list of roommates
	// in: body
	Body []RoommateDto
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
	Id string `json:"id"`
}

// swagger:parameters getRoommatesByUserId
//lint:ignore U1000 for docs
type roommatesUserIdParamsWrapper struct {
	// The user id
	// in: query
	// required: true
	UserId string `json:"userId"`
}

// swagger:parameters updateRoommate createRoommate
//lint:ignore U1000 for docs
type roommateParamsWrapper struct {
	// in: body
	// required: true
	Body RoommateDto
}

// swagger:parameters updateRoommate createRoommate getRoommatesByUserId getRoommate deleteRoommate
//lint:ignore U1000 for docs
type authHeaderWrapper struct {
	// in: header
	// required: true
	Token string
}
