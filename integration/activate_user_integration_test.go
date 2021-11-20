package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/user"
)

func activateUser(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := user.NewActivateUserParams()
	dto.ID = userId
	dto.Token = token
	_, err := cl.User.ActivateUser(dto)
	return err
}

func activateUserNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := user.NewActivateUserParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.User.ActivateUser(dto)
	if err != nil && !(strings.Contains(err.Error(), "activateUserForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
