package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/reference"
)

func getReference(t *testing.T, cl *client.OfTurnkeyAPI, referenceId string, token string) error {
	dto := reference.NewGetReferenceParams()
	dto.ID = referenceId
	dto.Token = token
	ok, err := cl.Reference.GetReference(dto)
	if err != nil {
		return err
	}
	if err = assert(ok.GetPayload().AdditionalDetails, "integration test AdditionalDetails", "AdditionalDetails"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Email, "integrationref@test.ca", "Email"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().FullName, "IntegrationRef Test", "FullName"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().PhoneNumber, "3068888888", "PhoneNumber"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Relationship, "integration test Relationship", "Relationship"); err != nil {
		return err
	}
	return nil
}

func getReferenceByUserId(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) (string, error) {
	dto := reference.NewGetReferencesByUserIDParams()
	dto.UserID = userId
	dto.Token = token
	ok, err := cl.Reference.GetReferencesByUserID(dto)
	if err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].AdditionalDetails, "integration test AdditionalDetails", "AdditionalDetails"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Email, "integrationref@test.ca", "Email"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].FullName, "IntegrationRef Test", "FullName"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].PhoneNumber, "3068888888", "PhoneNumber"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Relationship, "integration test Relationship", "Relationship"); err != nil {
		return "", err
	}
	return ok.GetPayload()[0].ID, nil
}

func getReferenceNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewGetReferenceParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Reference.GetReference(dto)
	if err != nil && !(strings.Contains(err.Error(), "getReferenceForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func getReferenceByUserIdNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := reference.NewGetReferencesByUserIDParams()
	dto.UserID = uuid.New().String()
	dto.Token = token
	_, err := cl.Reference.GetReferencesByUserID(dto)
	if err != nil && !(strings.Contains(err.Error(), "getReferencesByUserIdForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
