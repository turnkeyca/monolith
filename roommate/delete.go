package roommate

import (
	"fmt"
	"net/http"
)

// swagger:route DELETE /v1/roommate/{id} roommate deleteRoommate
// delete a roommate
//
// responses:
//	201: noContentResponse
//  404: roommateErrorResponse
//  500: roommateErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteRoommate(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	err := h.DeleteRoommate(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting roommate by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteRoommate(id string) error {
	err := h.db.Run("delete from roommate where id = $1;", id)
	return err
}
