package permission

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/authorizer"
	"github.com/turnkeyca/monolith/key"
)

// swagger:route POST /v1/permission permission createPermission
// create a new permission
//
// responses:
//	204: noContentResponse
//  400: permissionErrorResponse
//  422: permissionErrorResponse
//  500: permissionErrorResponse

// Create handles POST requests to add new permissions
func (h *Handler) HandlePostPermission(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*PermissionRequestDto)
	err := h.CreatePermissionRequest(dto)
	// TODO - RH - notify user somehow
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating permission request: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// swagger:route POST /v1/permission/{id}/accept permission acceptPermission
// accept a permission
//
// responses:
//	204: noContentResponse
//  500: permissionErrorResponse

// Accept handles POST requests to accept permission request
func (h *Handler) HandleAcceptPermission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.AcceptPermissionRequest(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error accepting permission request: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreatePermissionRequest(dto *PermissionRequestDto) error {
	id := uuid.New().String()
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
			$5
		);`,
		id,
		dto.UserId,
		dto.OnUserId,
		dto.PermissionRequestType,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}

func (h *Handler) AcceptPermissionRequest(id string) error {
	perm, err := h.GetPermission(id)
	if err != nil {
		return err
	}
	var newPerm authorizer.PermissionType
	if perm.Permission == authorizer.EDIT_PENDING {
		newPerm = authorizer.EDIT
	} else if perm.Permission == authorizer.VIEW_PENDING {
		newPerm = authorizer.VIEW
	} else {
		return fmt.Errorf("error accepting permission request: invalid permission type")
	}
	err = h.db.Run(`update "permission" set "permission"=$2, last_updated=$3 where id=$1`, id, newPerm, time.Now().Format(time.RFC3339Nano))
	return err
}
