package roommate

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/key"
)

// swagger:route POST /v1/roommate roommate createRoommate
// create a new roommate
//
// responses:
//	204: noContentResponse
//  400: roommateErrorResponse
//  422: roommateErrorResponse
//  500: roommateErrorResponse

// Create handles POST requests to add new roommates
func (h *Handler) HandlePostRoommate(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*RoommateDto)
	err := h.CreateRoommate(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating roommate: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateRoommate(dto *RoommateDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run(
		`insert into roommate (
			id, 
			user_id, 
			full_name, 
			email, 
			created_on,
			last_updated
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5,
			$5
		);`,
		dto.Id,
		dto.UserId,
		dto.FullName,
		dto.Email,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
