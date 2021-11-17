package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/user"
)

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
