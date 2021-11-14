package main

import (
	"fmt"
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/authentication"
	"github.com/turnkeyca/monolith/integration/models"
)

//RH - this function is too long on purpose.
func Test(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	transport := httptransport.New(fmt.Sprintf(`localhost:%s`, os.Getenv("PORT")), "", nil)
	// LOG IN
	cl := client.New(transport, strfmt.Default)
	userId, token, err := login(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	t.Logf(`user id: %s`, userId)
	t.Logf(`token: %s`, token)

	
}

func login(t *testing.T, cl *client.OfTurnkeyAPI) (string, string, error) {
	dto := authentication.NewRegisterNewTokenParams()
	newUserLoginId := uuid.New().String()
	t.Logf(`new user login id: %s`, newUserLoginId)
	dto.Body = &models.RegisterTokenDto{
		IsNewUser: true,
		LoginID:   newUserLoginId,
		Secret:    os.Getenv("SECRET_KEY"),
	}
	t.Logf(`body: %#v`, dto)
	ok, err := cl.Authentication.RegisterNewToken(dto)
	if err != nil {
		return "", "", err
	}
	return ok.GetPayload().ID, ok.GetPayload().Token, nil
}
