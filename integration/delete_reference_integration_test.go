package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/reference"
)

func deleteReference(t *testing.T, cl *client.OfTurnkeyAPI, referenceId string, token string) error {
	dto := reference.NewDeleteReferenceParams()
	dto.ID = referenceId
	dto.Token = token
	_, err := cl.Reference.DeleteReference(dto)
	return err
}

func deleteReferenceNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewDeleteReferenceParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Reference.DeleteReference(dto)
	if err != nil && !(strings.Contains(err.Error(), "deleteReferenceForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
