package permission

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// swagger:route POST /v1/permission permission createPermission
// create a new permission
//
// responses:
//	204: noContentResponse
//  422: permissionErrorValidation
//  500: permissionErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostPermission(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*PermissionDto)
	err := h.CreatePermission(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating permission: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreatePermission(dto *PermissionDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run(
		`insert into permission (
			id, 
			user_id, 
			on_user_id,
			permission,
			created_on, 
			last_updated
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6
		);`,
		dto.Id,
		dto.UserId,
		dto.OnUserId,
		dto.Permission,
		time.Now().Format(time.RFC3339Nano),
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
