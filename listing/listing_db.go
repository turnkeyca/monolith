package listing

import (
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

func (udb *ListingDatabase) SelectListing(id uuid.UUID) ([]Dto, error) {
	listings := []Dto{}
	err := udb.Select(&listings, "select * from listings where id = $1;", id.String())
	return listings, err
}
