package pet

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
)

type PetDto struct {
	Id          string `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId      string `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	PetType     string `json:"petType" db:"pet_type"`
	Breed       string `json:"breed" db:"breed"`
	SizeType    string `json:"sizeType" db:"size_type"`
	LastUpdated string `json:"lastUpdated" db:"last_updated"`
	CreatedOn   string `json:"createdOn" db:"created_on"`
}

func Read(r io.Reader) (*PetDto, error) {
	d := PetDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func WriteAll(e *[]PetDto, w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (d *PetDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *PetDto) Validate() error {
	v := validator.New()
	return v.Struct(d)
}
