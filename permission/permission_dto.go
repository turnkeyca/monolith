package permission

import (
	"encoding/json"
	"io"

	"github.com/go-playground/validator"
	"github.com/turnkeyca/monolith/util"
)

type PermissionType string
type PermissionRequestType string

const (
	VIEW         = "view"
	EDIT         = "edit"
	DECLINED     = "declined"
	VIEW_PENDING = "viewpending"
	EDIT_PENDING = "editpending"
)

var AllPermissionTypes []PermissionType = []PermissionType{VIEW, EDIT, DECLINED}
var AllPermissionRequestTypes []PermissionType = []PermissionType{VIEW_PENDING, EDIT_PENDING}

type PermissionRequestDto struct {
	UserId                string                `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	OnUserId              string                `json:"onUserId" validate:"omitempty,uuid" db:"on_user_id"`
	PermissionRequestType PermissionRequestType `json:"permission" validate:"permissionRequestType" db:"permission"`
}

type PermissionDto struct {
	Id          string         `json:"id" validate:"omitempty,uuid" db:"id"`
	UserId      string         `json:"userId" validate:"omitempty,uuid" db:"user_id"`
	OnUserId    string         `json:"onUserId" validate:"omitempty,uuid" db:"on_user_id"`
	Permission  PermissionType `json:"permission" validate:"permissionType" db:"permission"`
	LastUpdated string         `json:"lastUpdated" db:"last_updated"`
	CreatedOn   string         `json:"createdOn" db:"created_on"`
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
	return util.Contains(len(AllPermissionRequestTypes), func(index int) bool {
		return string(AllPermissionRequestTypes[index]) == fl.Field().String()
	})
}

func valPermissionType(fl validator.FieldLevel) bool {
	return util.Contains(len(AllPermissionTypes), func(index int) bool {
		return string(AllPermissionTypes[index]) == fl.Field().String()
	})
}
