package pet

import "github.com/turnkeyca/monolith/util"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response petErrorResponse
//lint:ignore U1000 for docs
type petErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body util.GenericError
}

// A pet
// swagger:response petResponse
//lint:ignore U1000 for docs
type petResponseWrapper struct {
	// A pet
	// in: body
	Body PetDto
}

// A list of pets
// swagger:response petsResponse
//lint:ignore U1000 for docs
type petsResponseWrapper struct {
	// A list of pets
	// in: body
	Body []PetDto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters updatePet getPet deletePet
//lint:ignore U1000 for docs
type petIdParamsWrapper struct {
	// The id of the pet for which the operation relates
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters getPetsByUserId
//lint:ignore U1000 for docs
type petsUserIdParamsWrapper struct {
	// The user id
	// in: query
	// required: true
	UserId string `json:"userId"`
}

// swagger:parameters updatePet createPet
//lint:ignore U1000 for docs
type petParamsWrapper struct {
	// in: body
	// required: true
	Body PetDto
}

// swagger:parameters updatePet getPet deletePet createPet getPetsByUserId
//lint:ignore U1000 for docs
type authHeaderWrapper struct {
	// in: header
	// required: true
	Token string
}
