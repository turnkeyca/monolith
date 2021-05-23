package pet

import (
	"os"

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

func (pdb *PetDatabase) SelectPet(id uuid.UUID) ([]Dto, error) {
	if os.Getenv("TEST") == "true" {
		pdb.PushQuery("select * from pet where id = $1;", id.String())
		dtos := []Dto{}
		for _, dto := range pdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(Dto))
		}
		return dtos, pdb.GetNextTestError()
	}
	pets := []Dto{}
	err := pdb.Select(&pets, "select * from pet where id = $1;", id.String())
	return pets, err
}
