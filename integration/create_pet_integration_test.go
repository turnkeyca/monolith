package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/pet"
	"github.com/turnkeyca/monolith/integration/models"
)

func createPet(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := pet.NewCreatePetParams()
	dto.Token = token
	dto.Body = &models.PetDto{
		UserID:   userId,
		Breed:    "Breed",
		PetType:  "PetType",
		SizeType: "SizeType",
	}
	_, err := cl.Pet.CreatePet(dto)
	return err
}

func createPetIncorrectUserId(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewCreatePetParams()
	dto.Token = token
	dto.Body = &models.PetDto{
		UserID:   uuid.New().String(),
		Breed:    "Breed",
		PetType:  "PetType",
		SizeType: "SizeType",
	}
	_, err := cl.Pet.CreatePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPetForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func createPetValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewCreatePetParams()
	dto.Token = token
	dto.Body = &models.PetDto{
		UserID:   "farts",
		Breed:    "Breed",
		PetType:  "PetType",
		SizeType: "SizeType",
	}
	_, err := cl.Pet.CreatePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPetUnprocessableEntity") && strings.Contains(err.Error(), "Error validating pet")) {
		return err
	}
	return nil
}
