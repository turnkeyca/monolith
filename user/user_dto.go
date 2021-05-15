package user

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Dto struct {
	Id                        uuid.UUID `json:"id" db:"id"`
	FullName                  string    `json:"fullName" validator:"required" db:"full_name"`
	Email                     string    `json:"email" validator:"required" db:"email"`
	Password                  string    `json:"password" db:"password"`
	PhoneNumber               string    `json:"phoneNumber" db:"phone_number"`
	Nickname                  string    `json:"nickname" validator:"required" db:"nickname"`
	Bio                       string    `json:"bio" validator:"required" db:"bio"`
	City                      string    `json:"city" validator:"required" db:"city"`
	Province                  string    `json:"province" validator:"required" db:"province"`
	UserType                  string    `json:"userType" validator:"required" db:"user_type"`
	SendNotifications         bool      `json:"sendNotifications" validator:"required" db:"send_notifications"`
	MovingReason              string    `json:"movingReason" db:"moving_reason"`
	HasRoommates              bool      `json:"roommates" validator:"required" db:"has_roommates"`
	HasSecurityDeposit        bool      `json:"securityDeposit" validator:"required" db:"has_security_deposit"`
	IsSmoker                  bool      `json:"smoker" validator:"required" db:"is_smoker"`
	HasPreviousLawsuit        bool      `json:"lawsuit" validator:"required" db:"has_prev_lawsuit"`
	HasPreviousEviction       bool      `json:"evicted" validator:"required" db:"has_prev_eviction"`
	CanCreditCheck            bool      `json:"creditCheck" validator:"required" db:"can_credit_check"`
	HasPets                   bool      `json:"pets" validator:"required" db:"has_pets"`
	AdditionalDetails         string    `json:"additionalDetails" db:"additional_details"`
	MoveInDate                string    `json:"moveInDate" db:"move_in_date"`
	MoveOutDate               string    `json:"moveOutDate" db:"move_out_date"`
	PropertyManagementCompany string    `json:"propertyManagementCompany" db:"property_management_company"`
	AdditionalDetailsLease    string    `json:"additionalDetailsLease" db:"additional_details_lease"`
	MonthlyBudgetMin          float64   `json:"monthlyBudgetMin" validator:"required" db:"monthly_budget_min"`
	MonthlyBudgetMax          float64   `json:"monthlyBudgetMax" validator:"required" db:"monthly_budget_max"`
}

func New() *Dto {
	return &Dto{
		Id:                        uuid.New(),
		FullName:                  "",
		Email:                     "",
		Password:                  "",
		PhoneNumber:               "",
		Nickname:                  "",
		Bio:                       "",
		City:                      "",
		Province:                  "",
		UserType:                  "",
		SendNotifications:         false,
		MovingReason:              "",
		HasRoommates:              false,
		HasSecurityDeposit:        false,
		IsSmoker:                  false,
		HasPreviousLawsuit:        false,
		HasPreviousEviction:       false,
		CanCreditCheck:            false,
		HasPets:                   false,
		AdditionalDetails:         "",
		MoveInDate:                "",
		MoveOutDate:               "",
		PropertyManagementCompany: "",
		AdditionalDetailsLease:    "",
		MonthlyBudgetMin:          0.00,
		MonthlyBudgetMax:          0.00,
	}
}

func Read(r io.Reader) (*Dto, error) {
	d := Dto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func (d *Dto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}

func (d *Dto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}
