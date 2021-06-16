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

func (rdb *ReferenceDatabase) SelectReference(id uuid.UUID) ([]ReferenceDto, error) {
	if os.Getenv("TEST") == "true" {
		rdb.PushQuery("select * from reference where id = $1;", id.String())
		dtos := []ReferenceDto{}
		for _, dto := range rdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(ReferenceDto))
		}
		return dtos, rdb.GetNextTestError()
	}
	references := []ReferenceDto{}
	err := rdb.Select(&references, "select * from reference where id = $1;", id.String())
	return references, err
}
