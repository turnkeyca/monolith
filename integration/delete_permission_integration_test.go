package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/permission"
)

func deletePermission(t *testing.T, cl *client.OfTurnkeyAPI, permissionId string, token string) error {
	dto := permission.NewDeletePermissionParams()
	dto.ID = permissionId
	dto.Token = token
	_, err := cl.Permission.DeletePermission(dto)
	return err
}

func deletePermissionNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := permission.NewDeletePermissionParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Permission.DeletePermission(dto)
	if err != nil && !(strings.Contains(err.Error(), "deletePermissionForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
