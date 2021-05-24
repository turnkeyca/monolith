package roommate

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type Dto struct {
	Id                string `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId            string `json:"userId" validate:"required,uuid" db:"user_id"`
	FullName          string `json:"fullName" validate:"required" db:"full_name"`
	Email             string `json:"email" validate:"omitempty,email" db:"email"`
	AdditionalDetails string `json:"additionalDetails" db:"additional_details"`
}

func Read(r io.Reader) (*Dto, error) {
	d := Dto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func (d *Dto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *Dto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}
