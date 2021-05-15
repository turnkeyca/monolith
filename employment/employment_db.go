package employment

import (
	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/db"
)

type EmploymentDatabase struct {
	*db.Database
}

func NewEmploymentDatabase(database *db.Database) *EmploymentDatabase {
	return &EmploymentDatabase{
		Database: database,
	}
}

func (udb *EmploymentDatabase) SelectEmployment(id uuid.UUID) ([]Dto, error) {
	employments := []Dto{}
	err := udb.Select(&employments, "select * from employments where id = $1;", id.String())
	return employments, err
}
