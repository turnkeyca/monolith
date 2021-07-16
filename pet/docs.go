package pet

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response petErrorResponse
//lint:ignore U1000 for docs
type petErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response petErrorValidation
//lint:ignore U1000 for docs
type petErrorValidationWrapper struct {
	// Collection of the validation errors
	// in: body
	Body ValidationError
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
