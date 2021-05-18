package listing

import (
	"os"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/db"
)

type ListingDatabase struct {
	*db.Database
}

func NewListingDatabase(database *db.Database) *ListingDatabase {
	return &ListingDatabase{
		Database: database,
	}
}

func (ldb *ListingDatabase) SelectListing(id uuid.UUID) ([]Dto, error) {
	if os.Getenv("TEST") == "true" {
		ldb.PushQuery("select * from employments where id = $1;", id.String())
		dtos := []Dto{}
		for _, dto := range ldb.GetNextTestReturn() {
			dtos = append(dtos, dto.(Dto))
		}
		return dtos, ldb.GetNextTestError()
	}
	listings := []Dto{}
	err := ldb.Select(&listings, "select * from listings where id = $1;", id.String())
	return listings, err
}
