package user

import (
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
	users := []Dto{}
	err := udb.Select(&users, "select * from users where id = $1;", id.String())
	return users, err
}
