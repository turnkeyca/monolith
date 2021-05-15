package user

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type Dto struct {
	Id      uuid.UUID  `json:"id" db:"id"`
	UserId  *uuid.UUID `json:"userId" db:"user_id"`
	Name    string     `json:"name" db:"name"`
	Address string     `json:"address" db:"address"`
	Link    string     `json:"link" db:"link"`
}

func New() *Dto {
	return &Dto{
		Id:      uuid.New(),
		UserId:  nil,
		Name:    "",
		Address: "",
		Link:    "",
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
