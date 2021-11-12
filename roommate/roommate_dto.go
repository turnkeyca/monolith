package roommate

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type RoommateDto struct {
	Id          string `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId      string `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	FullName    string `json:"fullName" db:"full_name"`
	Email       string `json:"email" db:"email"`
	LastUpdated string `json:"lastUpdated" db:"last_updated"`
	CreatedOn   string `json:"createdOn" db:"created_on"`
}

func Read(r io.Reader) (*RoommateDto, error) {
	d := RoommateDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func WriteAll(e *[]RoommateDto, w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (d *RoommateDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *RoommateDto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}
