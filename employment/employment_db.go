package employment

import (
	"os"

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

func (edb *EmploymentDatabase) SelectEmployment(id string) ([]EmploymentDto, error) {
	if os.Getenv("TEST") == "true" {
		edb.PushQuery("select * from employment where id = $1;", id)
		dtos := []EmploymentDto{}
		for _, dto := range edb.GetNextTestReturn() {
			dtos = append(dtos, dto.(EmploymentDto))
		}
		return dtos, edb.GetNextTestError()
	}
	employments := []EmploymentDto{}
	err := edb.Select(&employments, "select * from employment where id = $1;", id)
	return employments, err
}

func (edb *EmploymentDatabase) SelectEmploymentByUserId(id string) ([]EmploymentDto, error) {
	if os.Getenv("TEST") == "true" {
		edb.PushQuery("select * from employment where user_id = $1;", id)
		dtos := []EmploymentDto{}
		for _, dto := range edb.GetNextTestReturn() {
			dtos = append(dtos, dto.(EmploymentDto))
		}
		return dtos, edb.GetNextTestError()
	}
	employments := []EmploymentDto{}
	err := edb.Select(&employments, "select * from employment where user_id = $1;", id)
	return employments, err
}
