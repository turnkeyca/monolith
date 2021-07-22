package user

import (
	"testing"
)

func TestValidateLandlord(t *testing.T) {
	d := &UserDto{
		FullName:          "riley herman",
		Email:             "riley@herman.ca",
		Password:          "password",
		PhoneNumber:       "3069999999",
		Nickname:          "nick",
		UserType:          LANDLORD,
		SendNotifications: true,
	}
	err := d.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

func TestValidateRenter(t *testing.T) {
	d := &UserDto{
		FullName:          "riley herman",
		Email:             "riley@herman.ca",
		Password:          "password",
		PhoneNumber:       "3069999999",
		Nickname:          "nick",
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
