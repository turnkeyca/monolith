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
		Breed:    "integration test Breed update",
		PetType:  "integratioon test PetType update",
		SizeType: "integration test SizeType update",
	}
	_, err := cl.Pet.UpdatePet(dto)
	return err
}

func updatePetNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewUpdatePetParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	dto.Body = &models.PetDto{
		Breed:    "integration test Breed update",
		PetType:  "integratioon test PetType update",
		SizeType: "integration test SizeType update",
	}
	_, err := cl.Pet.UpdatePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPetForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func updatePetValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewUpdatePetParams()
	dto.Token = token
	dto.Body = &models.PetDto{
		UserID:   "farts",
		Breed:    "integration test Breed update",
		PetType:  "integratioon test PetType update",
		SizeType: "integration test SizeType update",
	}
	_, err := cl.Pet.UpdatePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPetForbidden") && strings.Contains(err.Error(), "Error reading pet")) {
		return err
	}
	return nil
}
