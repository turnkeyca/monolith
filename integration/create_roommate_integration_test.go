package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/roommate"
	"github.com/turnkeyca/monolith/integration/models"
)

func createRoommate(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := roommate.NewCreateRoommateParams()
	dto.Token = token
	dto.Body = &models.RoommateDto{
		UserID:   userId,
		Email:    "integrationref@test.ca",
		FullName: "IntegrationRef Test",
	}
	_, err := cl.Roommate.CreateRoommate(dto)
	return err
}

func createRoommateIncorrectUserId(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := roommate.NewCreateRoommateParams()
	dto.Token = token
	dto.Body = &models.RoommateDto{
		UserID:   uuid.New().String(),
		Email:    "integrationref@test.ca",
		FullName: "IntegrationRef Test",
	}
	_, err := cl.Roommate.CreateRoommate(dto)
	if err != nil && !(strings.Contains(err.Error(), "createRoommateForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func createRoommateValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := roommate.NewCreateRoommateParams()
	dto.Token = token
	dto.Body = &models.RoommateDto{
		UserID:   "farts",
		Email:    "integrationref@test.ca",
		FullName: "IntegrationRef Test",
	}
	_, err := cl.Roommate.CreateRoommate(dto)
	if err != nil && !(strings.Contains(err.Error(), "createRoommateUnprocessableEntity") && strings.Contains(err.Error(), "Error validating roommate")) {
		return err
	}
	return nil
}
