package roommate

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route POST /api/roommate roommate createRoommate
// create a new roommate
//
// responses:
//	200: roommateResponse
//  422: roommateErrorValidation
//  500: roommateErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostRoommate(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*Dto)
	err := h.CreateRoommate(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating roommate: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateRoommate(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into roommate (id, full_name) values ($1, $2);", dto.Id.String(), dto.FullName)
	return err
}
