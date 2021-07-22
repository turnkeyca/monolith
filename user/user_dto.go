package user

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/turnkeyca/monolith/util"
)

type UserType string

const (
	RENTER   = "renter"
	LANDLORD = "landlord"
)

var AllUserTypes []UserType = []UserType{RENTER, LANDLORD}

type UserStatusType string

const (
	INACTIVE = "inactive"
	ACTIVE   = "active"
)

var AllUserStatusTypes []UserStatusType = []UserStatusType{INACTIVE, ACTIVE}

type UserDto struct {
	Id                       string         `json:"id" db:"id"`
	FullName                 string         `json:"fullName" db:"full_name"`
	Email                    string         `json:"email" db:"email"`
	Password                 string         `json:"password" db:"password"`
	UserStatus               UserStatusType `json:"userStatusType" validate:"omitempty,userStatusType" db:"user_status"`
	LastUpdated              string         `json:"lastUpdated" db:"last_updated"`
	CreatedOn                string         `json:"createdOn" db:"created_on"`
	PhoneNumber              string         `json:"phoneNumber" db:"phone_number"`
	Nickname                 string         `json:"nickname" db:"nickname"`
	Bio                      string         `json:"bio" db:"bio"`
	UserType                 UserType       `json:"userType" validate:"omitempty,userType" db:"user_type"`
	SendNotifications        bool           `json:"sendNotifications" db:"send_notifications"`
	MovingReason             string         `json:"movingReason" db:"moving_reason"`
	HasRoommates             bool           `json:"roommates" db:"has_roommates"`
	HasSecurityDeposit       bool           `json:"securityDeposit" db:"has_security_deposit"`
	IsSmoker                 bool           `json:"smoker" db:"is_smoker"`
	HasPreviousLawsuit       bool           `json:"lawsuit" db:"has_prev_lawsuit"`
	HasPreviousEviction      bool           `json:"evicted" db:"has_prev_eviction"`
	CanCreditCheck           bool           `json:"creditCheck" db:"can_credit_check"`
	HasPets                  bool           `json:"pets" db:"has_pets"`
	AdditionalDetailsGeneral string         `json:"additionalDetailsGeneral" db:"additional_details"`
	MoveInDate               string         `json:"moveInDate" db:"move_in_date"`
	MoveOutDate              string         `json:"moveOutDate" db:"move_out_date"`
	AdditionalDetailsLease   string         `json:"additionalDetailsLease" db:"additional_details_lease"`
}

func Read(r io.Reader) (*UserDto, error) {
	d := UserDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func (d *UserDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *UserDto) Validate() error {
	v := validator.New()
	v.RegisterValidation("renterRequired", valRenterRequired)
	v.RegisterValidation("renter", valRenter)
	v.RegisterValidation("landlord", valLandlord)
	v.RegisterValidation("userType", valUserType)
	v.RegisterValidation("userStatusType", valUserStatusType)
	return v.Struct(d)
}

func valRenterRequired(fl validator.FieldLevel) bool {
	d := fl.Parent().Interface().(*UserDto)
	if d.UserType == RENTER {
		return fl.Field().String() != ""
	}
	return true
}

func valRenter(fl validator.FieldLevel) bool {
	d := fl.Parent().Interface().(*UserDto)
	if d.UserType != RENTER {
		return fl.Field().IsZero()
	}
	return true
}

func valLandlord(fl validator.FieldLevel) bool {
	d := fl.Parent().Interface().(*UserDto)
	if d.UserType != LANDLORD {
		return fl.Field().IsZero()
	}
	return true
}

func valUserType(fl validator.FieldLevel) bool {
	return util.Contains(len(AllUserTypes), func(index int) bool {
		return string(AllUserTypes[index]) == fl.Field().String()
	})
}

func valUserStatusType(fl validator.FieldLevel) bool {
	return util.Contains(len(AllUserStatusTypes), func(index int) bool {
		return string(AllUserStatusTypes[index]) == fl.Field().String()
	})
}
