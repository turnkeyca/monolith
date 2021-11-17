package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/reference"
	"github.com/turnkeyca/monolith/integration/models"
)

func createReference(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := reference.NewCreateReferenceParams()
	dto.Token = token
	dto.Body = &models.ReferenceDto{
		UserID:            userId,
		AdditionalDetails: "integration test AdditionalDetails",
		Email:             "integrationref@test.ca",
		FullName:          "IntegrationRef Test",
		PhoneNumber:       "3068888888",
		Relationship:      "integration test Relationship",
	}
	_, err := cl.Reference.CreateReference(dto)
	return err
}

func createReferenceIncorrectUserId(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewCreateReferenceParams()
	dto.Token = token
	dto.Body = &models.ReferenceDto{
		UserID:            uuid.New().String(),
		AdditionalDetails: "integration test AdditionalDetails",
		Email:             "integrationref@test.ca",
		FullName:          "IntegrationRef Test",
		PhoneNumber:       "3068888888",
		Relationship:      "integration test Relationship",
	}
	_, err := cl.Reference.CreateReference(dto)
	if err != nil && !(strings.Contains(err.Error(), "createReferenceForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func createReferenceValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewCreateReferenceParams()
	dto.Token = token
	dto.Body = &models.ReferenceDto{
		UserID:            "farts",
		AdditionalDetails: "integration test AdditionalDetails",
		Email:             "integrationref@test.ca",
		FullName:          "IntegrationRef Test",
		PhoneNumber:       "3068888888",
		Relationship:      "integration test Relationship",
	}
	_, err := cl.Reference.CreateReference(dto)
	if err != nil && !(strings.Contains(err.Error(), "createReferenceUnprocessableEntity") && strings.Contains(err.Error(), "Error validating reference")) {
		return err
	}
	return nil
}
