package roommate

import (
	"os"

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

func (rdb *RoommateDatabase) SelectRoommate(id uuid.UUID) ([]RoommateDto, error) {
	if os.Getenv("TEST") == "true" {
		rdb.PushQuery("select * from roommate where id = $1;", id.String())
		dtos := []RoommateDto{}
		for _, dto := range rdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(RoommateDto))
		}
		return dtos, rdb.GetNextTestError()
	}
	roommates := []RoommateDto{}
	err := rdb.Select(&roommates, "select * from roommate where id = $1;", id.String())
	return roommates, err
}
