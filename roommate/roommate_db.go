package roommate

import (
	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/db"
)

type RoommateDatabase struct {
	*db.Database
}

func NewRoommateDatabase(database *db.Database) *RoommateDatabase {
	return &RoommateDatabase{
		Database: database,
	}
}

func (udb *RoommateDatabase) SelectRoommate(id uuid.UUID) ([]Dto, error) {
	roommates := []Dto{}
	err := udb.Select(&roommates, "select * from roommates where id = $1;", id.String())
	return roommates, err
}
