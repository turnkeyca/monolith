package pet

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type PetDto struct {
	Id     string  `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId string  `json:"userId" validate:"required,uuid" db:"user_id"`
	Breed  string  `json:"breed" db:"breed"`
	Weight float64 `json:"weight" db:"weight"`
}

func Read(r io.Reader) (*PetDto, error) {
	d := PetDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func (d *PetDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *PetDto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}
