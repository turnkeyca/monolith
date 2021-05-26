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
	dto := r.Context().Value(KeyBody{}).(*RoommateDto)
	err := h.CreateRoommate(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating roommate: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateRoommate(dto *RoommateDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run("insert into roommate (id, user_id, full_name, email, additional_details) values ($1, $2, $3, $4, $5);", dto.Id, dto.UserId, dto.FullName, dto.Email, dto.AdditionalDetails)
	return err
}
