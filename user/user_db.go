package user

import (
	"os"

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

func (udb *UserDatabase) SelectUser(id string) ([]UserDto, error) {
	if os.Getenv("TEST") == "true" {
		udb.PushQuery("select * from employments where id = $1;", id)
		dtos := []UserDto{}
		for _, dto := range udb.GetNextTestReturn() {
			dtos = append(dtos, dto.(UserDto))
		}
		return dtos, udb.GetNextTestError()
	}
	users := []UserDto{}
	err := udb.Select(&users, "select * from users where id = $1;", id)
	return users, err
}
