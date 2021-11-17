package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/roommate"
	"github.com/turnkeyca/monolith/integration/models"
)

func updateRoommate(t *testing.T, cl *client.OfTurnkeyAPI, roommateId string, token string) error {
	dto := roommate.NewUpdateRoommateParams()
	dto.ID = roommateId
	dto.Token = token
	dto.Body = &models.RoommateDto{
		Email:    "integrationrefupdate@test.ca",
		FullName: "IntegrationRefUpdate Test",
	}
	_, err := cl.Roommate.UpdateRoommate(dto)
	return err
}

func updateRoommateNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := roommate.NewUpdateRoommateParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	dto.Body = &models.RoommateDto{
		Email:    "integrationrefupdate@test.ca",
		FullName: "IntegrationRefUpdate Test",
	}
	_, err := cl.Roommate.UpdateRoommate(dto)
	if err != nil && !(strings.Contains(err.Error(), "updateRoommateNotFound") && strings.Contains(err.Error(), "not found")) {
		return err
	}
	return nil
}

func updateRoommateValidationError(t *testing.T, cl *client.OfTurnkeyAPI, roommateId string, token string) error {
	dto := roommate.NewUpdateRoommateParams()
	dto.Token = token
	dto.ID = roommateId
	dto.Body = &models.RoommateDto{
		UserID:   "farts",
		Email:    "integrationrefupdate@test.ca",
		FullName: "IntegrationRefUpdate Test",
	}
	_, err := cl.Roommate.UpdateRoommate(dto)
	if err != nil && !(strings.Contains(err.Error(), "updateRoommateUnprocessableEntity") && strings.Contains(err.Error(), "Error validating roommate")) {
		return err
	}
	return nil
}
