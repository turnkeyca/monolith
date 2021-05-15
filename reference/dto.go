package user

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Dto struct {
	Id                uuid.UUID  `json:"id" db:"id"`
	UserId            *uuid.UUID `json:"userId" db:"user_id"`
	FullName          string     `json:"fullName" validator:"required" db:"full_name"`
	Email             string     `json:"email" db:"email"`
	PhoneNumber       string     `json:"phoneNumber" db:"phone_number"`
	Relationship      string     `json:"relationship" validator:"required" db:"relationship"`
	AdditionalDetails string     `json:"additionalDetails" db:"additional_details"`
}

func New() *Dto {
	return &Dto{
		Id:                uuid.New(),
		UserId:            nil,
		FullName:          "",
		Email:             "",
		PhoneNumber:       "",
		Relationship:      "",
		AdditionalDetails: "",
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
