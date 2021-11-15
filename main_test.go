package main

import (
	"fmt"
	"os"
	"strings"
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
	err = loginBadSecretShouldFail(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = loginNewUserNewUserFlagFalse(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	userId, token, err := login(t, cl)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = loginReturningUserNewUserFlag(t, cl, userId)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}

	// DELETE, UPDATE, GET USER
	err = deleteUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = deleteUserNotFoundShouldNotError(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateUserValidationError(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = updateUserNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getUser(t, cl, userId, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
	err = getUserNotFound(t, cl, token)
	if err != nil {
		t.Logf(`error: %s`, err)
		t.Fail()
	}
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

func deleteUser(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
	dto := user.NewDeleteUserParams()
	dto.ID = userId
	dto.Token = token
	_, err := cl.User.DeleteUser(dto)
	return err
}

func deleteUserNotFoundShouldNotError(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := user.NewDeleteUserParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.User.DeleteUser(dto)
	if err != nil && !(strings.Contains(err.Error(), "deleteUserForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
		return err
	}
	return nil
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

func updateUserNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := user.NewUpdateUserParams()
	dto.ID = uuid.New().String()
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
	if err != nil && !(strings.Contains(err.Error(), "updateUserForbidden") && strings.Contains(err.Error(), "")) {
		return err
	}
	return nil
}

func updateUserValidationError(t *testing.T, cl *client.OfTurnkeyAPI, userId string, token string) error {
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
		UserType:                 "gobbledegook",
	}
	_, err := cl.User.UpdateUser(dto)
	if err != nil && !(strings.Contains(err.Error(), "updateUserUnprocessableEntity") && strings.Contains(err.Error(), "validation for 'UserType' failed on the 'userType' tag")) {
		return err
	}
	return nil
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

func getUserNotFound(t *testing.T, cl *client.OfTurnkeyAPI, token string) error {
	dto := user.NewGetUserParams()
	dto.ID = uuid.New().String()
	dto.Token = token
	_, err := cl.User.GetUser(dto)
	if err != nil && !(strings.Contains(err.Error(), "getUserForbidden") && strings.Contains(err.Error(), "User does not have permission")) {
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
