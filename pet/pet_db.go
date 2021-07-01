package pet

import (
	"os"

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

func (pdb *PetDatabase) SelectPet(id string) ([]PetDto, error) {
	if os.Getenv("TEST") == "true" {
		pdb.PushQuery("select * from pet where id = $1;", id)
		dtos := []PetDto{}
		for _, dto := range pdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(PetDto))
		}
		return dtos, pdb.GetNextTestError()
	}
	pets := []PetDto{}
	err := pdb.Select(&pets, "select * from pet where id = $1;", id)
	return pets, err
}
