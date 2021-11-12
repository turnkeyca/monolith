package permission

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/turnkeyca/monolith/authorizer"
	"github.com/turnkeyca/monolith/util"
)

type PermissionRequestDto struct {
	UserId                string                           `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	OnUserId              string                           `json:"onUserId" validate:"omitempty,uuid" db:"on_user_id"`
	PermissionRequestType authorizer.PermissionRequestType `json:"permission" validate:"permissionRequestType" db:"permission"`
}

type PermissionDto struct {
	Id          string                    `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId      string                    `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	OnUserId    string                    `json:"onUserId" validate:"omitempty,uuid" db:"on_user_id"`
	Permission  authorizer.PermissionType `json:"permission" validate:"permissionType" db:"permission"`
	LastUpdated string                    `json:"lastUpdated" db:"last_updated"`
	CreatedOn   string                    `json:"createdOn" db:"created_on"`
}

func ReadPermissionRequestDto(r io.Reader) (*PermissionRequestDto, error) {
	d := PermissionRequestDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func Read(r io.Reader) (*PermissionDto, error) {
	d := PermissionDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}

func WriteAll(e *[]PermissionDto, w io.Writer) error {
	return json.NewEncoder(w).Encode(e)
}

func (d *PermissionDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}

func (d *PermissionRequestDto) Validate() error {
	v := validator.New()
	v.RegisterValidation("permissionRequestType", valPermissionRequestType)
	return v.Struct(d)
}

func (d *PermissionDto) Validate() error {
	v := validator.New()
	v.RegisterValidation("permissionType", valPermissionType)
	return v.Struct(d)
}

func valPermissionRequestType(fl validator.FieldLevel) bool {
	return util.Contains(len(authorizer.AllPermissionRequestTypes), func(index int) bool {
		return string(authorizer.AllPermissionRequestTypes[index]) == fl.Field().String()
	})
}

func valPermissionType(fl validator.FieldLevel) bool {
	return util.Contains(len(authorizer.AllPermissionTypes), func(index int) bool {
		return string(authorizer.AllPermissionTypes[index]) == fl.Field().String()
	})
}
