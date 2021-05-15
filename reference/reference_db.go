package reference

import (
	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/db"
)

type ReferenceDatabase struct {
	*db.Database
}

func NewReferenceDatabase(database *db.Database) *ReferenceDatabase {
	return &ReferenceDatabase{
		Database: database,
	}
}

func (udb *ReferenceDatabase) SelectReference(id uuid.UUID) ([]Dto, error) {
	references := []Dto{}
	err := udb.Select(&references, "select * from references where id = $1;", id.String())
	return references, err
}
