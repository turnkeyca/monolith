package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/roommate"
)

func deleteRoommate(t *testing.T, cl *client.OfTurnkeyAPI, roommateId string, token string) error {
	dto := roommate.NewDeleteRoommateParams()
	dto.ID = roommateId
	dto.Token = token
	_, err := cl.Roommate.DeleteRoommate(dto)
	return err
}

func deleteRoommateNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := roommate.NewDeleteRoommateParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Roommate.DeleteRoommate(dto)
	if err != nil && !(strings.Contains(err.Error(), "deleteRoommateForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
