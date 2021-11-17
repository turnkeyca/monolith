package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/permission"
	"github.com/turnkeyca/monolith/integration/models"
)

func createPermission(t *testing.T, cl *client.OfTurnkeyAPI, userId string, onUserId string, token string) error {
	dto := permission.NewCreatePermissionParams()
	dto.Token = token
	dto.Body = &models.PermissionDto{
		UserID:     userId,
		OnUserID:   onUserId,
		Permission: "viewpending",
	}
	_, err := cl.Permission.CreatePermission(dto)
	return err
}

func createPermissionIncorrectUserId(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := permission.NewCreatePermissionParams()
	dto.Token = token
	dto.Body = &models.PermissionDto{
		UserID:     uuid.New().String(),
		OnUserID:   uuid.New().String(),
		Permission: "viewpending",
	}
	_, err := cl.Permission.CreatePermission(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPermissionForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func createPermissionUserIdIsOnUserId(t *testing.T, cl *client.OfTurnkeyAPI, userId string, onUserId string, token string) error {
	dto := permission.NewCreatePermissionParams()
	dto.Token = token
	dto.Body = &models.PermissionDto{
		UserID:     onUserId,
		OnUserID:   userId,
		Permission: "viewpending",
	}
	_, err := cl.Permission.CreatePermission(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPermissionForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func createPermissionValidationError(t *testing.T, cl *client.OfTurnkeyAPI, userId string, onUserId string, token string) error {
	dto := permission.NewCreatePermissionParams()
	dto.Token = token
	dto.Body = &models.PermissionDto{
		UserID:     userId,
		OnUserID:   onUserId,
		Permission: "farts",
	}
	_, err := cl.Permission.CreatePermission(dto)
	if err != nil && !(strings.Contains(err.Error(), "createPermissionUnprocessableEntity") && strings.Contains(err.Error(), "Error validating permission")) {
		return err
	}
	return nil
}
