package employment

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type EmploymentDto struct {
	Id                string  `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId            string  `json:"userId" validate:"required,uuid" db:"user_id"`
	Employer          string  `json:"employer" validate:"required" db:"employer"`
	Occupation        string  `json:"occupation" validate:"required" db:"occupation"`
	Duration          string  `json:"duration" validate:"required" db:"duration"`
	AdditionalDetails string  `json:"additionalDetails" db:"additional_details"`
	AnnualSalary      float64 `json:"annualSalary" db:"annual_salary"`
}

func Read(r io.Reader) (*EmploymentDto, error) {
	d := EmploymentDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func WriteAll(e *[]EmploymentDto, w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (d *EmploymentDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *EmploymentDto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}
