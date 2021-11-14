package main

import (
	"fmt"
	"os"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/authentication"
	"github.com/turnkeyca/monolith/integration/models"
)

func TestLogin(t *testing.T) {
	err := godotenv.Load(".env")
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	transport := httptransport.New(fmt.Sprintf(`localhost:%s`, os.Getenv("PORT")), "", nil)
	cl := client.New(transport, strfmt.Default)
	dto := authentication.NewRegisterNewTokenParams()
	dto.Body = &models.RegisterTokenDto{
		IsNewUser: true,
		LoginID:   "grandpariley",
		Secret:    os.Getenv("SECRET_KEY"),
	}
	t.Logf(`body: %#v`, dto)
	ok, err := cl.Authentication.RegisterNewToken(dto)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	t.Logf(`ok: %#v`, ok)
	userId := ok.GetPayload().ID
	t.Logf(`user id: %s`, userId)
	token := ok.GetPayload().Token
	t.Logf(`token: %s`, token)
}
