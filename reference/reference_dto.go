package reference

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type ReferenceDto struct {
	Id                string `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId            string `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	FullName          string `json:"fullName" db:"full_name"`
	Email             string `json:"email" db:"email"`
	PhoneNumber       string `json:"phoneNumber" db:"phone_number"`
	Relationship      string `json:"relationship" db:"relationship"`
	AdditionalDetails string `json:"additionalDetails" db:"additional_details"`
	LastUpdated       string `json:"lastUpdated" db:"last_updated"`
	CreatedOn         string `json:"createdOn" db:"created_on"`
}

func Read(r io.Reader) (*ReferenceDto, error) {
	d := ReferenceDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func WriteAll(e *[]ReferenceDto, w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (d *ReferenceDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *ReferenceDto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}
