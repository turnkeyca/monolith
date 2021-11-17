package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/pet"
)

func deletePet(t *testing.T, cl *client.OfTurnkeyAPI, petId string, token string) error {
	dto := pet.NewDeletePetParams()
	dto.ID = petId
	dto.Token = token
	_, err := cl.Pet.DeletePet(dto)
	return err
}

func deletePetNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewDeletePetParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Pet.DeletePet(dto)
	if err != nil && !(strings.Contains(err.Error(), "deletePetNotFound") && strings.Contains(err.Error(), "not found")) {
		return err
	}
	return nil
}
