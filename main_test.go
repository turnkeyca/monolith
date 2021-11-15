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
	"github.com/turnkeyca/monolith/integration/client/user"
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
	err = updateUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
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

func updateUser(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := user.NewUpdateUserParams()
	dto.ID = userId
	dto.Token = token
	dto.Body = &models.UserDto{
		AcceptedTerms:            true,
		AdditionalDetailsGeneral: "integration test AdditionalDetailsGeneral",
		AdditionalDetailsLease:   "integration test AdditionalDetailsLease",
		Bio:                      "integration test Bio",
		CanCreditCheck:           true,
		Email:                    "integration@test.ca",
		FullName:                 "Integration Test",
		HasPets:                  true,
		HasPreviousEviction:      true,
		HasPreviousLawsuit:       true,
		HasRoommates:             true,
		HasSecurityDeposit:       true,
		IsSmoker:                 true,
		MoveInDate:               "2021-04-13",
		MoveOutDate:              "2022-04-13",
		MovingReason:             "integration test MovingReason",
		Nickname:                 "testy",
		PhoneNumber:              "3069999999",
		SendNotifications:        true,
		WalkthroughComplete:      false,
		UserType:                 "renter",
	}
	_, err := cl.User.UpdateUser(dto)
	return err
}

func getUser(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := user.NewGetUserParams()
	dto.ID = userId
	dto.Token = token
	ok, err := cl.User.GetUser(dto)
	if err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().AcceptedTerms, "AcceptedTerms"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().AdditionalDetailsGeneral, "integration test AdditionalDetailsGeneral", "AdditionalDetailsGeneral"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().AdditionalDetailsLease, "integration test AdditionalDetailsLease", "AdditionalDetailsLease"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Bio, "integration test Bio", "Bio"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().CanCreditCheck, "CanCreditCheck"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Email, "integration@test.ca", "Email"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().FullName, "Integration Test", "FullName"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().HasPets, "HasPets"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().HasPreviousEviction, "HasPreviousEviction"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().HasPreviousLawsuit, "HasPreviousLawsuit"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().HasRoommates, "HasRoommates"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().HasSecurityDeposit, "HasSecurityDeposit"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().IsSmoker, "IsSmoker"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().MoveInDate, "2021-04-13", "MoveInDate"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().MoveOutDate, "2022-04-13", "MoveOutDate"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().MovingReason, "integration test MovingReason", "MovingReason"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().Nickname, "testy", "Nickname"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().PhoneNumber, "3069999999", "PhoneNumber"); err != nil {
		return err
	}
	if err = assertTrue(ok.GetPayload().SendNotifications, "SendNotifications"); err != nil {
		return err
	}
	if err = assertTrue(!ok.GetPayload().WalkthroughComplete, "WalkthroughComplete"); err != nil {
		return err
	}
	if err = assert(string(ok.GetPayload().UserType), "renter", "UserType"); err != nil {
		return err
	}
	if err = assert(ok.GetPayload().ID, userId, "Id"); err != nil {
		return err
	}
	return nil
}

func assert(obj1 string, obj2 string, name string) error {
	if obj1 != obj2 {
		return fmt.Errorf(`assert failed: expected %s to be %s but was %s`, name, obj1, obj2)
	}
	return nil
}

func assertTrue(obj bool, name string) error {
	if !obj {
		return fmt.Errorf(`assert failed: expected %s to be %t but was %t`, name, obj, !obj)
	}
	return nil
}
