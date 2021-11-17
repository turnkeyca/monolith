package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/pet"
	"github.com/turnkeyca/monolith/integration/models"
)

func updatePet(t *testing.T, cl *client.OfTurnkeyAPI, petId string, token string) error {
	dto := pet.NewUpdatePetParams()
	dto.ID = petId
	dto.Token = token
	dto.Body = &models.PetDto{
		Breed:    "updateB",
		PetType:  "updateP",
		SizeType: "updateS",
	}
	_, err := cl.Pet.UpdatePet(dto)
	return err
}

func updatePetNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewUpdatePetParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	dto.Body = &models.PetDto{
		Breed:    "updateB",
		PetType:  "updateP",
		SizeType: "updateS",
	}
	_, err := cl.Pet.UpdatePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "updatePetNotFound") && strings.Contains(err.Error(), "not found")) {
		return err
	}
	return nil
}

func updatePetValidationError(t *testing.T, cl *client.OfTurnkeyAPI, petId string, token string) error {
	dto := pet.NewUpdatePetParams()
	dto.Token = token
	dto.ID = petId
	dto.Body = &models.PetDto{
		UserID:   "farts",
		Breed:    "updateB",
		PetType:  "updateP",
		SizeType: "updateS",
	}
	_, err := cl.Pet.UpdatePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "updatePetUnprocessableEntity") && strings.Contains(err.Error(), "Error validating pet")) {
		return err
	}
	return nil
}
