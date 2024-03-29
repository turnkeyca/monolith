package permission

import "github.com/turnkeyca/monolith/util"

//
// NOTE: Types defined here are purely for documentation purposes
// these types are not used by any of the handlers

// Generic error message returned as a string
// swagger:response permissionErrorResponse
//lint:ignore U1000 for docs
type permissionErrorResponseWrapper struct {
	// Description of the error
	// in: body
	Body util.GenericError
}

// A permission
// swagger:response permissionResponse
//lint:ignore U1000 for docs
type permissionResponseWrapper struct {
	// A permission
	// in: body
	Body PermissionDto
}

// A list of permissions
// swagger:response permissionsResponse
//lint:ignore U1000 for docs
type permissionsResponseWrapper struct {
	// A list of permissions
	// in: body
	Body []PermissionDto
}

// No content is returned by this API endpoint
// swagger:response noContentResponse
//lint:ignore U1000 for docs
type noContentResponseWrapper struct {
}

// swagger:parameters acceptPermission getPermission deletePermission
//lint:ignore U1000 for docs
type permissionIdParamsWrapper struct {
	// The id of the permission for which the operation relates
	// in: path
	// required: true
	Id string `json:"id"`
}

// swagger:parameters getPermissionsByUserId
//lint:ignore U1000 for docs
type permissionsUserIdParamsWrapper struct {
	// The user id
	// in: query
	// required: true
	UserId string `json:"userId"`
}

// swagger:parameters  createPermission
//lint:ignore U1000 for docs
type permissionParamsWrapper struct {
	// in: body
	// required: true
	Body PermissionDto
}

// swagger:parameters acceptPermission createPermission getPermissionsByUserId getPermission deletePermission
//lint:ignore U1000 for docs
type authHeaderWrapper struct {
	// in: header
	// required: true
	Token string
}
