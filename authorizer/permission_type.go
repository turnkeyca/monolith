package authorizer

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
