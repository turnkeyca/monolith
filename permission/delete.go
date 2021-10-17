package permission

import (
	"fmt"
	"net/http"
)

// swagger:route DELETE /v1/permission/{id} permission deletePermission
// delete a permission
//
// responses:
//	201: noContentResponse
//  404: permissionErrorResponse
//  500: permissionErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeletePermission(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	err := h.DeletePermission(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting permission by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeletePermission(id string) error {
	err := h.db.Run("delete from permission where id = $1;", id)
	return err
}
