package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/employment"
	"github.com/turnkeyca/monolith/integration/models"
)

func updateEmployment(t *testing.T, cl *client.OfTurnkeyAPI, employmentId string, token string) error {
	dto := employment.NewUpdateEmploymentParams()
	dto.ID = employmentId
	dto.Token = token
	dto.Body = &models.EmploymentDto{
		AdditionalDetails: "integration test AdditionalDetails update",
		Duration:          "Duration update",
		Employer:          "integration test Employer update",
		Occupation:        "integration test Occupation update",
		RentAffordability: "integration test RentAffordability update",
		AnnualSalary:      100001,
	}
	_, err := cl.Employment.UpdateEmployment(dto)
	return err
}

func updateEmploymentNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewUpdateEmploymentParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	dto.Body = &models.EmploymentDto{
		AdditionalDetails: "integration test AdditionalDetails update",
		Duration:          "Duration update",
		Employer:          "integration test Employer update",
		Occupation:        "integration test Occupation update",
		RentAffordability: "integration test RentAffordability update",
		AnnualSalary:      100001,
	}
	_, err := cl.Employment.UpdateEmployment(dto)
	if err != nil && !(strings.Contains(err.Error(), "createEmploymentForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func updateEmploymentValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewUpdateEmploymentParams()
	dto.Token = token
	dto.Body = &models.EmploymentDto{
		UserID:            "farts",
		AdditionalDetails: "integration test AdditionalDetails update",
		Duration:          "Duration update",
		Employer:          "integration test Employer update",
		Occupation:        "integration test Occupation update",
		RentAffordability: "integration test RentAffordability update",
		AnnualSalary:      100001,
	}
	_, err := cl.Employment.UpdateEmployment(dto)
	if err != nil && !(strings.Contains(err.Error(), "createEmploymentForbidden") && strings.Contains(err.Error(), "Error reading employment")) {
		return err
	}
	return nil
}
