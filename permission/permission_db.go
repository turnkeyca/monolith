package permission

import (
	"os"

	"github.com/turnkeyca/monolith/db"
)

type PermissionDatabase struct {
	*db.Database
}

func NewPermissionDatabase(database *db.Database) *PermissionDatabase {
	return &PermissionDatabase{
		Database: database,
	}
}

func (rdb *PermissionDatabase) SelectPermission(id string) ([]PermissionDto, error) {
	if os.Getenv("TEST") == "true" {
		rdb.PushQuery("select * from permission where id = $1;", id)
		dtos := []PermissionDto{}
		for _, dto := range rdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(PermissionDto))
		}
		return dtos, rdb.GetNextTestError()
	}
	permissions := []PermissionDto{}
	err := rdb.Select(&permissions, "select * from permission where id = $1;", id)
	return permissions, err
}

func (rdb *PermissionDatabase) SelectPermissionsByUserId(id string) ([]PermissionDto, error) {
	if os.Getenv("TEST") == "true" {
		rdb.PushQuery("select * from permission where user_id = $1;", id)
		dtos := []PermissionDto{}
		for _, dto := range rdb.GetNextTestReturn() {
			dtos = append(dtos, dto.(PermissionDto))
		}
		return dtos, rdb.GetNextTestError()
	}
	permissions := []PermissionDto{}
	err := rdb.Select(&permissions, "select * from permission where user_id = $1;", id)
	return permissions, err
}
