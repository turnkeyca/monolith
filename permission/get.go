package permission

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route GET /v1/permission/{id} permission getPermission
// return a permission
// responses:
//	200: permissionResponse
//  403: permissionErrorResponse
//  404: permissionErrorResponse
//	500: permissionErrorResponse

// HandleGetPermission handles GET requests
func (h *Handler) HandleGetPermission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	permission, err := h.GetPermission(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting permission by id: %s, %s", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = permission.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

// swagger:route GET /v1/permission permission getPermissionsByUserId
// return all permissions for a user
// responses:
//	200: permissionsResponse
//  403: permissionErrorResponse
//	500: permissionErrorResponse

// HandleGetPermissionByUserId handles GET requests
func (h *Handler) HandleGetPermissionByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyUserId{}).(string)
	permissions, err := h.GetPermissionByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting permission by user id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(permissions, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetPermission(id string) (*PermissionDto, error) {
	var permissions []PermissionDto
	err := h.db.Select(&permissions, `select * from permission where id = $1;`, id)
	if err != nil {
		return nil, err
	}
	if permissions == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(permissions) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &permissions[0], err
}

func (h *Handler) GetPermissionByUserId(userId string) (*[]PermissionDto, error) {
	var permissions []PermissionDto
	err := h.db.Select(&permissions, `select * from permission where user_id = $1 or on_user_id = $1;`, userId)
	if err != nil {
		return nil, err
	}
	if permissions == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &permissions, err
}
