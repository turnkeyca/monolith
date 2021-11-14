package permission

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/authorizer"
	"github.com/turnkeyca/monolith/key"
)

// swagger:route DELETE /v1/permission/{id} permission deletePermission
// delete a permission
//
// responses:
//	204: noContentResponse
//  409: permissionErrorResponse
//  500: permissionErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeletePermission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.DeletePermission(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting permission by id: %s, %s", id, err), http.StatusConflict)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeletePermission(id string) error {
	perm, err := h.GetPermission(id)
	if err != nil {
		return err
	}
	if perm.OnUserId == perm.UserId {
		return fmt.Errorf("cannot delete base permission")
	}
	return h.db.Run(`update "permission" set "permission"=$2, last_updated=$3 where id=$1;`, id, authorizer.DECLINED, time.Now().Format(time.RFC3339Nano))
}
