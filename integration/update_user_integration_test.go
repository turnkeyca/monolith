package integration

import (
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/integration/client"
	"github.com/turnkeyca/monolith/integration/client/user"
	"github.com/turnkeyca/monolith/integration/models"
)

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
