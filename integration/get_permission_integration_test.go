package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/permission"
	"github.com/turnkeyca/monolith/integration/models"
)

func getPermission(t *testing.T, cl *client.OfTurnkeyAPI, permissionId string, userId string, onUserId string, token string) error {
	dto := permission.NewGetPermissionParams()
	dto.ID = permissionId
	dto.Token = token
	ok, err := cl.Permission.GetPermission(dto)
	if err != nil {
		return err
	}
	if err = assert(ok.GetPayload().UserID, userId, "UserID"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().OnUserID, onUserId, "OnUserID"); err != nil {
		return err
	}
	if err = assert(string(ok.GetPayload().Permission), "view", "Permission"); err != nil {
		return err
	}
	return nil
}

func getPermissionByUserId(t *testing.T, cl *client.OfTurnkeyAPI, userId string, onUserId string, token string) (string, error) {
	dto := permission.NewGetPermissionsByUserIDParams()
	dto.UserID = userId
	dto.Token = token
	ok, err := cl.Permission.GetPermissionsByUserID(dto)
	var perm *models.PermissionDto
	for _, payload := range ok.GetPayload() {
		if payload.OnUserID != payload.UserID {
			perm = payload
		}
	}
	if err != nil {
		return "", err
	}
	if err = assert(perm.UserID, userId, "UserID"); err != nil {
		return "", err
	}
	if err = assert(perm.OnUserID, onUserId, "OnUserID"); err != nil {
		return "", err
	}
	if err = assert(string(perm.Permission), "view", "Permission"); err != nil {
		return "", err
	}
	return perm.ID, nil
}

func getPermissionNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := permission.NewGetPermissionParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Permission.GetPermission(dto)
	if err != nil && !(strings.Contains(err.Error(), "getPermissionForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func getPermissionByUserIdNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := permission.NewGetPermissionsByUserIDParams()
	dto.UserID = uuid.New().String()
	dto.Token = token
	_, err := cl.Permission.GetPermissionsByUserID(dto)
	if err != nil && !(strings.Contains(err.Error(), "getPermissionsByUserIdForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
