package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/user"
)

func deleteUser(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := user.NewDeleteUserParams()
	dto.ID = userId
	dto.Token = token
	_, err := cl.User.DeleteUser(dto)
	return err
}

func deleteUserNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := user.NewDeleteUserParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.User.DeleteUser(dto)
	if err != nil && !(strings.Contains(err.Error(), "deleteUserForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
