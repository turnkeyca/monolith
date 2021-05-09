package user

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Dto struct {
	Id       uuid.UUID `json:"id" db:"id"`
	FullName string    `json:"fullName" validator:"required" db:"full_name"`
}

func New() *Dto {
	return &Dto{
		Id:       uuid.New(),
		FullName: "",
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
