package roommate

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route GET /api/roommate/{id} roommate getRoommate
// return a roommate
// responses:
//	200: roommateResponse
//	404: roommateErrorResponse

// HandleGetRoommate handles GET requests
func (h *Handler) HandleGetRoommate(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	roommate, err := h.GetRoommate(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting roommate by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = roommate.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetRoommate(id uuid.UUID) (*RoommateDto, error) {
	result, err := NewRoommateDatabase(h.db).SelectRoommate(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for id: %s", id.String())
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id.String())
	}
	return &result[0], err
}
