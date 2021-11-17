package roommate

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route DELETE /v1/roommate/{id} roommate deleteRoommate
// delete a roommate
//
// responses:
//	204: noContentResponse
//  403: roommateErrorResponse
//  404: roommateErrorResponse
//  500: roommateErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteRoommate(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.DeleteRoommate(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting roommate by id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteRoommate(id string) error {
	return h.db.Run(`delete from roommate where id = $1;`, id)
}
