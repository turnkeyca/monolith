package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/permission"
)

func acceptPermission(t *testing.T, cl *client.OfTurnkeyAPI, permissionId string, token string) error {
	dto := permission.NewAcceptPermissionParams()
	dto.ID = permissionId
	dto.Token = token
	_, err := cl.Permission.AcceptPermission(dto)
	return err
}

func acceptPermissionNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := permission.NewAcceptPermissionParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Permission.AcceptPermission(dto)
	if err != nil && !(strings.Contains(err.Error(), "acceptPermissionForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
