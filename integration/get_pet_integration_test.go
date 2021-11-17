package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/pet"
)

func getPet(t *testing.T, cl *client.OfTurnkeyAPI, petId string, token string) error {
	dto := pet.NewGetPetParams()
	dto.ID = petId
	dto.Token = token
	ok, err := cl.Pet.GetPet(dto)
	if err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Breed, "Breed", "Breed"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().PetType, "PetType", "PetType"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().SizeType, "SizeType", "SizeType"); err != nil {
		return err
	}
	return nil
}

func getPetByUserId(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) (string, error) {
	dto := pet.NewGetPetsByUserIDParams()
	dto.UserID = userId
	dto.Token = token
	ok, err := cl.Pet.GetPetsByUserID(dto)
	if err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Breed, "Breed", "Breed"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].PetType, "PetType", "PetType"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].SizeType, "SizeType", "SizeType"); err != nil {
		return "", err
	}
	return ok.GetPayload()[0].ID, nil
}

func getPetNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewGetPetParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Pet.GetPet(dto)
	if err != nil && !(strings.Contains(err.Error(), "getPetNotFound") && strings.Contains(err.Error(), "not found")) {
		return err
	}
	return nil
}

func getPetByUserIdNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := pet.NewGetPetsByUserIDParams()
	dto.UserID = uuid.New().String()
	dto.Token = token
	_, err := cl.Pet.GetPetsByUserID(dto)
	if err != nil && !(strings.Contains(err.Error(), "getPetsByUserIdForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
