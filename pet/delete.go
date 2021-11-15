package pet

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route DELETE /v1/pet/{id} pet deletePet
// delete a pet
//
// responses:
//	204: noContentResponse
//  403: petErrorResponse
//  500: petErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeletePet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.DeletePet(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting pet by id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeletePet(id string) error {
	return h.db.Run(`delete from pet where id = $1`, id)
}
