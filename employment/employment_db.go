package employment

import (
	"os"

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

func (edb *EmploymentDatabase) SelectEmployment(id uuid.UUID) ([]Dto, error) {
	if os.Getenv("TEST") == "true" {
		edb.PushQuery("select * from employments where id = $1;", id.String())
		dtos := []Dto{}
		for _, dto := range edb.GetNextTestReturn() {
			dtos = append(dtos, dto.(Dto))
		}
		return dtos, edb.GetNextTestError()
	}
	employments := []Dto{}
	err := edb.Select(&employments, "select * from employments where id = $1;", id.String())
	return employments, err
}
