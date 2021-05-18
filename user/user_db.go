package user

import (
	"os"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/db"
)

type UserDatabase struct {
	*db.Database
}

func NewUserDatabase(database *db.Database) *UserDatabase {
	return &UserDatabase{
		Database: database,
	}
}

func (udb *UserDatabase) SelectUser(id uuid.UUID) ([]Dto, error) {
	if os.Getenv("TEST") == "true" {
		udb.PushQuery("select * from employments where id = $1;", id.String())
		dtos := []Dto{}
		for _, dto := range udb.GetNextTestReturn() {
			dtos = append(dtos, dto.(Dto))
		}
		return dtos, udb.GetNextTestError()
	}
	users := []Dto{}
	err := udb.Select(&users, "select * from users where id = $1;", id.String())
	return users, err
}
