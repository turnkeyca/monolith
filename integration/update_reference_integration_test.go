package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/reference"
	"github.com/turnkeyca/monolith/integration/models"
)

func updateReference(t *testing.T, cl *client.OfTurnkeyAPI, referenceId string, token string) error {
	dto := reference.NewUpdateReferenceParams()
	dto.ID = referenceId
	dto.Token = token
	dto.Body = &models.ReferenceDto{
		AdditionalDetails: "integration test AdditionalDetails update",
		Email:             "integrationrefupdate@test.ca",
		FullName:          "IntegrationRefUpdate Test",
		PhoneNumber:       "3068888887",
		Relationship:      "integration test Relationship update",
	}
	_, err := cl.Reference.UpdateReference(dto)
	return err
}

func updateReferenceNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewUpdateReferenceParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	dto.Body = &models.ReferenceDto{
		AdditionalDetails: "integration test AdditionalDetails update",
		Email:             "integrationrefupdate@test.ca",
		FullName:          "IntegrationRefUpdate Test",
		PhoneNumber:       "3068888887",
		Relationship:      "integration test Relationship update",
	}
	_, err := cl.Reference.UpdateReference(dto)
	if err != nil && !(strings.Contains(err.Error(), "createReferenceForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func updateReferenceValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewUpdateReferenceParams()
	dto.Token = token
	dto.Body = &models.ReferenceDto{
		UserID:            "farts",
		AdditionalDetails: "integration test AdditionalDetails update",
		Email:             "integrationrefupdate@test.ca",
		FullName:          "IntegrationRefUpdate Test",
		PhoneNumber:       "3068888887",
		Relationship:      "integration test Relationship update",
	}
	_, err := cl.Reference.UpdateReference(dto)
	if err != nil && !(strings.Contains(err.Error(), "createReferenceForbidden") && strings.Contains(err.Error(), "Error reading reference")) {
		return err
	}
	return nil
}
