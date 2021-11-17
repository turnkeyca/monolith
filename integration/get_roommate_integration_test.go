package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/roommate"
)

func getRoommate(t *testing.T, cl *client.OfTurnkeyAPI, roommateId string, token string) error {
	dto := roommate.NewGetRoommateParams()
	dto.ID = roommateId
	dto.Token = token
	ok, err := cl.Roommate.GetRoommate(dto)
	if err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Email, "integrationref@test.ca", "Email"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().FullName, "IntegrationRef Test", "FullName"); err != nil {
		return err
	}
	return nil
}

func getRoommateByUserId(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) (string, error) {
	dto := roommate.NewGetRoommatesByUserIDParams()
	dto.UserID = userId
	dto.Token = token
	ok, err := cl.Roommate.GetRoommatesByUserID(dto)
	if err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Email, "integrationref@test.ca", "Email"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].FullName, "IntegrationRef Test", "FullName"); err != nil {
		return "", err
	}
	return ok.GetPayload()[0].ID, nil
}

func getRoommateNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := roommate.NewGetRoommateParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Roommate.GetRoommate(dto)
	if err != nil && !(strings.Contains(err.Error(), "getRoommateNotFound") && strings.Contains(err.Error(), "not found")) {
		return err
	}
	return nil
}

func getRoommateByUserIdNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := roommate.NewGetRoommatesByUserIDParams()
	dto.UserID = uuid.New().String()
	dto.Token = token
	_, err := cl.Roommate.GetRoommatesByUserID(dto)
	if err != nil && !(strings.Contains(err.Error(), "getRoommatesByUserIdForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
