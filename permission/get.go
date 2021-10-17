package permission

import (
	"fmt"
	"net/http"
)

// swagger:route GET /v1/permission/{id} permission getPermission
// return a permission
// responses:
//	200: permissionResponse
//	404: permissionErrorResponse

// HandleGetPermission handles GET requests
func (h *Handler) HandleGetPermission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	permission, err := h.GetPermission(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting permission by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = permission.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %#v", err), http.StatusInternalServerError)
	}
}

// swagger:route GET /v1/permission permission getPermissionsByUserId
// return all permissions ofr a user
// responses:
//	200: permissionsResponse
//	404: permissionErrorResponse

// HandleGetPermissionByUserId handles GET requests
func (h *Handler) HandleGetPermissionByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyUserId{}).(string)
	permissions, err := h.GetPermissionByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting permission by user id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(permissions, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %#v", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetPermission(id string) (*PermissionDto, error) {
	result, err := NewPermissionDatabase(h.db).SelectPermission(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &result[0], err
}

func (h *Handler) GetPermissionByUserId(userId string) (*[]PermissionDto, error) {
	result, err := NewPermissionDatabase(h.db).SelectPermissionsByUserId(userId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &result, err
}
