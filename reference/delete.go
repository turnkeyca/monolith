package reference

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route DELETE /api/reference/{id} reference deleteReference
// delete a reference
//
// responses:
//	201: noContentResponse
//  404: referenceErrorResponse
//  500: referenceErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteReference(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	err := h.DeleteReference(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting reference by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteReference(id uuid.UUID) error {
	err := h.db.Run("delete from references where id = $1;", id.String())
	return err
}
