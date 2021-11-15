package integration

import (
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/authentication"
	"github.com/turnkeyca/monolith/integration/models"
)

func loginBadSecretShouldFail(t *testing.T, cl *client.OfTurnkeyAPI) error {
	dto := authentication.NewRegisterNewTokenParams()
	newUserLoginId := uuid.New().String()
	dto.Body = &models.RegisterTokenDto{
		IsNewUser: true,
		LoginID:   newUserLoginId,
		Secret:    uuid.New().String(),
	}
	_, err := cl.Authentication.RegisterNewToken(dto)
	if err != nil && !(strings.Contains(err.Error(), "registerNewTokenInternalServerError") && strings.Contains(err.Error(), "invalid secret key")) {
		return err
	}
	return nil
}

func login(t *testing.T, cl *client.OfTurnkeyAPI) (string, string, error) {
	dto := authentication.NewRegisterNewTokenParams()
	newUserLoginId := uuid.New().String()
	dto.Body = &models.RegisterTokenDto{
		IsNewUser: true,
		LoginID:   newUserLoginId,
		Secret:    os.Getenv("SECRET_KEY"),
	}
	ok, err := cl.Authentication.RegisterNewToken(dto)
	if err != nil {
		return "", "", err
	}
	return ok.GetPayload().ID, ok.GetPayload().Token, nil
}

func loginNewUserNewUserFlagFalse(t *testing.T, cl *client.OfTurnkeyAPI) error {
	dto := authentication.NewRegisterNewTokenParams()
	newUserLoginId := uuid.New().String()
	dto.Body = &models.RegisterTokenDto{
		IsNewUser: false,
		LoginID:   newUserLoginId,
		Secret:    os.Getenv("SECRET_KEY"),
	}
	_, err := cl.Authentication.RegisterNewToken(dto)
	if err != nil && !(strings.Contains(err.Error(), "registerNewTokenInternalServerError") && strings.Contains(err.Error(), "user does not exist")) {
		return err
	}
	return nil
}

func loginReturningUserNewUserFlag(t *testing.T, cl *client.OfTurnkeyAPI, userId string) error {
	dto := authentication.NewRegisterNewTokenParams()
	dto.Body = &models.RegisterTokenDto{
		IsNewUser: true,
		LoginID:   userId,
		Secret:    os.Getenv("SECRET_KEY"),
	}
	_, err := cl.Authentication.RegisterNewToken(dto)
	if err != nil && !(strings.Contains(err.Error(), "registerNewTokenInternalServerError") && strings.Contains(err.Error(), "duplicate user")) {
		return err
	}
	return nil
}
