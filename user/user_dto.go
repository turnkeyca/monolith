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

type ProvinceType string

const (
	YUKON                 = "YK"
	NORTHWEST_TERRITORIES = "NW"
	NUNAVUT               = "NU"
	BRITISH_COLUMBIA      = "BC"
	ALBERTA               = "AB"
	SASKATCHEWAN          = "SK"
	MANITOBA              = "MB"
	ONTARIO               = "ON"
	QUEBEC                = "QC"
	NEWFOUNDLAND_LABRADOR = "NL"
	NEW_BRUNSWICK         = "NB"
	NOVA_SCOTIA           = "NS"
	PRINCE_EDWARD_ISLAND  = "PE"
)

var AllProvinceTypes []ProvinceType = []ProvinceType{YUKON, NORTHWEST_TERRITORIES, NUNAVUT, BRITISH_COLUMBIA, ALBERTA, SASKATCHEWAN, MANITOBA, ONTARIO, QUEBEC, NEWFOUNDLAND_LABRADOR, NEW_BRUNSWICK, NOVA_SCOTIA, PRINCE_EDWARD_ISLAND}

type UserDto struct {
	Id                        string       `json:"id" validate:"omitempty,uuid" db:"id"`
	FullName                  string       `json:"fullName" validate:"required" db:"full_name"`
	Email                     string       `json:"email" validate:"required,email" db:"email"`
	Password                  string       `json:"password" validate:"required" db:"password"`
	PhoneNumber               string       `json:"phoneNumber" db:"phone_number"`
	Nickname                  string       `json:"nickname" validate:"required" db:"nickname"`
	Bio                       string       `json:"bio" db:"bio"`
	City                      string       `json:"city" validate:"required" db:"city"`
	Province                  ProvinceType `json:"province" validate:"required,provinceType" db:"province"`
	UserType                  UserType     `json:"userType" validate:"required,userType" db:"user_type"`
	SendNotifications         bool         `json:"sendNotifications" validate:"required" db:"send_notifications"`
	MovingReason              string       `json:"movingReason" validate:"renter,renterRequired" db:"moving_reason"`
	HasRoommates              bool         `json:"roommates" validate:"renter,renterRequired" db:"has_roommates"`
	HasSecurityDeposit        bool         `json:"securityDeposit" validate:"renter,renterRequired" db:"has_security_deposit"`
	IsSmoker                  bool         `json:"smoker" validate:"renter,renterRequired" db:"is_smoker"`
	HasPreviousLawsuit        bool         `json:"lawsuit" validate:"renter,renterRequired" db:"has_prev_lawsuit"`
	HasPreviousEviction       bool         `json:"evicted" validate:"renter,renterRequired" db:"has_prev_eviction"`
	CanCreditCheck            bool         `json:"creditCheck" validate:"renter,renterRequired" db:"can_credit_check"`
	HasPets                   bool         `json:"pets" validate:"renter,renterRequired" db:"has_pets"`
	AdditionalDetails         string       `json:"additionalDetails" validate:"renter" db:"additional_details"`
	MoveInDate                string       `json:"moveInDate" validate:"renter,renterRequired" db:"move_in_date"`
	MoveOutDate               string       `json:"moveOutDate" validate:"renter,renterRequired" db:"move_out_date"`
	PropertyManagementCompany string       `json:"propertyManagementCompany" validate:"landlord" db:"property_management_company"`
	AdditionalDetailsLease    string       `json:"additionalDetailsLease" validate:"renter" db:"additional_details_lease"`
	MonthlyBudgetMin          float64      `json:"monthlyBudgetMin" validate:"renter,renterRequired" db:"monthly_budget_min"`
	MonthlyBudgetMax          float64      `json:"monthlyBudgetMax" validate:"renter,renterRequired" db:"monthly_budget_max"`
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
	v.RegisterValidation("provinceType", valProvinceType)
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

func valProvinceType(fl validator.FieldLevel) bool {
	return util.Contains(len(AllProvinceTypes), func(index int) bool {
		return string(AllProvinceTypes[index]) == fl.Field().String()
	})
}
