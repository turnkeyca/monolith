package permission

import (
	"fmt"
	"net/http"
	"time"
)

// swagger:route PUT /v1/permission/{id} permission updatePermission
// update a permission
//
// responses:
//	201: noContentResponse
//  404: permissionErrorResponse
//  422: permissionErrorValidation

// Update handles PUT requests to update permissions
func (h *Handler) HandlePutPermission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	dto := r.Context().Value(KeyBody{}).(*PermissionDto)
	dto.Id = id
	err := h.UpdatePermission(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating permission: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdatePermission(dto *PermissionDto) error {
	err := h.db.Run(
		`update permission set 
			id=$1, 
			user_id=$2, 
			on_user_id=$3, 
			permission=$4, 
			last_updated=$5
		where id=$1;`,
		dto.Id,
		dto.UserId,
		dto.OnUserId,
		dto.Permission,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
