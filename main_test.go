package main

import (
	"os"
	"testing"

	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/auth"
	"github.com/turnkeyca/monolith/integration/models"
)

func TestLogin(t *testing.T) {
	dto := &auth.RegisterNewTokenParams{
		Body: &models.RegisterTokenDto{
			IsNewUser: true,
			LoginID:   "grandpariley",
			Secret:    os.Getenv("SECRET"),
		},
	}
	cl := client.NewHTTPClient(nil)
	ok, err := cl.Auth.RegisterNewToken(dto)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	userId := ok.GetPayload().ID
	t.Logf(`user id: %s`, userId)
	token := ok.GetPayload().Token
	t.Logf(`token: %s`, token)
}
