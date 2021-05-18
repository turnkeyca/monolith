package reference

import (
	"os"

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

func (rdb *ReferenceDatabase) SelectReference(id uuid.UUID) ([]Dto, error) {
	if os.Getenv("TEST") == "true" {
		rdb.PushQuery("select * from employments where id = $1;", id.String())
		dtos := []Dto{}
		for _, dto := range rdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(Dto))
		}
		return dtos, rdb.GetNextTestError()
	}
	references := []Dto{}
	err := rdb.Select(&references, "select * from references where id = $1;", id.String())
	return references, err
}
