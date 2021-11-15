package integration

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/employment"
)

func getEmployment(t *testing.T, cl *client.OfTurnkeyAPI, employmentId string, token string) error {
	dto := employment.NewGetEmploymentParams()
	dto.ID = employmentId
	dto.Token = token
	ok, err := cl.Employment.GetEmployment(dto)
	if err != nil {
		return err
	}
	if err = assert(ok.GetPayload().AdditionalDetails, "integration test AdditionalDetails", "AdditionalDetails"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Duration, "Duration", "Duration"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Employer, "integration test Employer", "Employer"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Occupation, "integration test Occupation", "Occupation"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().RentAffordability, "integration test RentAffordability", "RentAffordability"); err != nil {
		return err
	}
	if err = assert(strconv.FormatFloat(ok.GetPayload().AnnualSalary, 'f', 2, 64), "100000.00", "AnnualSalary"); err != nil {
		return err
	}
	return nil
}

func getEmploymentByUserId(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) (string, error) {
	dto := employment.NewGetEmploymentsByUserIDParams()
	dto.UserID = userId
	dto.Token = token
	ok, err := cl.Employment.GetEmploymentsByUserID(dto)
	if err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].AdditionalDetails, "integration test AdditionalDetails", "AdditionalDetails"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Duration, "Duration", "Duration"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Employer, "integration test Employer", "Employer"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].Occupation, "integration test Occupation", "Occupation"); err != nil {
		return "", err
	}
	if err = assert(ok.GetPayload()[0].RentAffordability, "integration test RentAffordability", "RentAffordability"); err != nil {
		return "", err
	}
	if err = assert(strconv.FormatFloat(ok.GetPayload()[0].AnnualSalary, 'f', 2, 64), "100000.00", "AnnualSalary"); err != nil {
		return "", err
	}
	return ok.GetPayload()[0].ID, nil
}

func getEmploymentNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewGetEmploymentParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.Employment.GetEmployment(dto)
	if err != nil && !(strings.Contains(err.Error(), "getEmploymentForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}

func getEmploymentByUserIdNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := employment.NewGetEmploymentsByUserIDParams()
	dto.UserID = uuid.New().String()
	dto.Token = token
	_, err := cl.Employment.GetEmploymentsByUserID(dto)
	if err != nil && !(strings.Contains(err.Error(), "getEmploymentsByUserIdForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
}
