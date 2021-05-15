package employment

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Dto struct {
	Id                uuid.UUID  `json:"id" db:"id"`
	EmploymentId      *uuid.UUID `json:"employmentId" validator:"required" db:"employment_id"`
	Employer          string     `json:"employer" validator:"required" db:"employer"`
	Occupation        string     `json:"occupation" validator:"required" db:"occupation"`
	Duration          string     `json:"duration" validator:"required" db:"duration"`
	AdditionalDetails string     `json:"additionalDetails" db:"additional_details"`
	AnnualSalary      float64    `json:"annualSalary" db:"annual_salary"`
}

func New() *Dto {
	return &Dto{
		Id:                uuid.New(),
		EmploymentId:      nil,
		Occupation:        "",
		Duration:          "",
		AdditionalDetails: "",
		AnnualSalary:      0.00,
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
