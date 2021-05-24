package user

import (
	"testing"
)

func TestValidateLandlord(t *testing.T) {
	d := &Dto{
		FullName:                  "riley herman",
		Email:                     "riley@herman.ca",
		Password:                  "password",
		PhoneNumber:               "3069999999",
		Nickname:                  "nick",
		City:                      "Regina",
		Province:                  SASKATCHEWAN,
		UserType:                  LANDLORD,
		SendNotifications:         true,
		PropertyManagementCompany: "PMC Co",
	}
	err := d.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateRenter(t *testing.T) {
	d := &Dto{
		FullName:          "riley herman",
		Email:             "riley@herman.ca",
		Password:          "password",
		PhoneNumber:       "3069999999",
		Nickname:          "nick",
		City:              "Regina",
		Province:          SASKATCHEWAN,
		UserType:          RENTER,
		SendNotifications: true,
		MovingReason:      "because I can",
		MoveInDate:        "2021-05-13",
		MoveOutDate:       "2021-05-13",
	}
	err := d.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
