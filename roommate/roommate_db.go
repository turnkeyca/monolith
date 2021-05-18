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

func (rdb *RoommateDatabase) SelectRoommate(id uuid.UUID) ([]Dto, error) {
	if os.Getenv("TEST") == "true" {
		rdb.PushQuery("select * from employments where id = $1;", id.String())
		dtos := []Dto{}
		for _, dto := range rdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(Dto))
		}
		return dtos, rdb.GetNextTestError()
	}
	roommates := []Dto{}
	err := rdb.Select(&roommates, "select * from roommates where id = $1;", id.String())
	return roommates, err
}
