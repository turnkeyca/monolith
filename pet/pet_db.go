package pet

import (
	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/db"
)

type PetDatabase struct {
	*db.Database
}

func NewPetDatabase(database *db.Database) *PetDatabase {
	return &PetDatabase{
		Database: database,
	}
}

func (udb *PetDatabase) SelectPet(id uuid.UUID) ([]Dto, error) {
	pets := []Dto{}
	err := udb.Select(&pets, "select * from pets where id = $1;", id.String())
	return pets, err
}
