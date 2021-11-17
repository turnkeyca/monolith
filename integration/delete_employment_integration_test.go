package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/employment"
)

func deleteEmployment(t *testing.T, cl *client.OfTurnkeyAPI, employmentId string, token string) error {
	dto := employment.NewDeleteEmploymentParams()
	dto.ID = employmentId
	dto.Token = token
	_, err := cl.Employment.DeleteEmployment(dto)
	return err
}

func deleteEmploymentNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewDeleteEmploymentParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Employment.DeleteEmployment(dto)
	if err != nil && !(strings.Contains(err.Error(), "deleteEmploymentNotFound") && strings.Contains(err.Error(), "not found")) {
		return err
	}
	return nil
}
