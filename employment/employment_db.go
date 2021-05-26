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

func (edb *EmploymentDatabase) SelectEmployment(id uuid.UUID) ([]EmploymentDto, error) {
	if os.Getenv("TEST") == "true" {
		edb.PushQuery("select * from employment where id = $1;", id.String())
		dtos := []EmploymentDto{}
		for _, dto := range edb.GetNextTestReturn() {
			dtos = append(dtos, dto.(EmploymentDto))
		}
		return dtos, edb.GetNextTestError()
	}
	employments := []EmploymentDto{}
	err := edb.Select(&employments, "select * from employment where id = $1;", id.String())
	return employments, err
}
