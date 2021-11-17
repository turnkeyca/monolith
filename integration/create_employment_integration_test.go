package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/employment"
	"github.com/turnkeyca/monolith/integration/models"
)

func createEmployment(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := employment.NewCreateEmploymentParams()
	dto.Token = token
	dto.Body = &models.EmploymentDto{
		UserID:            userId,
		AdditionalDetails: "integration test AdditionalDetails",
		Duration:          "Duration",
		Employer:          "integration test Employer",
		Occupation:        "integration test Occupation",
		RentAffordability: "integration test RentAffordability",
		AnnualSalary:      100000,
	}
	_, err := cl.Employment.CreateEmployment(dto)
	return err
}

func createEmploymentIncorrectUserId(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewCreateEmploymentParams()
	dto.Token = token
	dto.Body = &models.EmploymentDto{
		UserID:            uuid.New().String(),
		AdditionalDetails: "integration test AdditionalDetails",
		Duration:          "Duration",
		Employer:          "integration test Employer",
		Occupation:        "integration test Occupation",
		RentAffordability: "integration test RentAffordability",
		AnnualSalary:      100000,
	}
	_, err := cl.Employment.CreateEmployment(dto)
	if err != nil && !(strings.Contains(err.Error(), "createEmploymentForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func createEmploymentValidationError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewCreateEmploymentParams()
	dto.Token = token
	dto.Body = &models.EmploymentDto{
		UserID:            "farts",
		AdditionalDetails: "integration test AdditionalDetails",
		Duration:          "Duration",
		Employer:          "integration test Employer",
		Occupation:        "integration test Occupation",
		RentAffordability: "integration test RentAffordability",
		AnnualSalary:      100000,
	}
	_, err := cl.Employment.CreateEmployment(dto)
	if err != nil && !(strings.Contains(err.Error(), "createEmploymentUnprocessableEntity") && strings.Contains(err.Error(), "Error validating employment")) {
		return err
	}
	return nil
}
