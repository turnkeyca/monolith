package pet

import (
	"fmt"
	"net/http"
)

// swagger:route DELETE /v1/pet/{id} pet deletePet
// delete a pet
//
// responses:
//	201: noContentResponse
//  404: petErrorResponse
//  500: petErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeletePet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	err := h.DeletePet(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting pet by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeletePet(id string) error {
	err := h.db.Run("delete from pet where id = $1;", id)
	return err
}
