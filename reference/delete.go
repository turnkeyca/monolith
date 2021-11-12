package reference

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route DELETE /v1/reference/{id} reference deleteReference
// delete a reference
//
// responses:
//	204: noContentResponse
//  500: referenceErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteReference(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.DeleteReference(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting reference by id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteReference(id string) error {
	return h.db.Run(`delete from reference where id = $1;`, id)
}
