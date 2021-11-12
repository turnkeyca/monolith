package roommate

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route GET /v1/roommate/{id} roommate getRoommate
// return a roommate
// responses:
//	200: roommateResponse
//	404: roommateErrorResponse
//	500: roommateErrorResponse

// HandleGetRoommate handles GET requests
func (h *Handler) HandleGetRoommate(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	roommate, err := h.GetRoommate(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting roommate by id: %s, %s", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = roommate.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

// swagger:route GET /v1/roommate roommate getRoommatesByUserId
// return all roommates for a user
// responses:
//	200: roommatesResponse
//  500: roommateErrorResponse

// HandleGetRoommateByUserId handles GET requests
func (h *Handler) HandleGetRoommateByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyUserId{}).(string)
	roommates, err := h.GetRoommateByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting roommate by user id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(roommates, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetRoommate(id string) (*RoommateDto, error) {
	var roommates []RoommateDto
	err := h.db.Select(&roommates, `select * from roommate where id = $1;`, id)
	if err != nil {
		return nil, err
	}
	if roommates == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(roommates) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &roommates[0], err
}

func (h *Handler) GetRoommateByUserId(userId string) (*[]RoommateDto, error) {
	var roommates []RoommateDto
	err := h.db.Select(&roommates, `select * from roommate where user_id = $1;`, userId)
	if err != nil {
		return nil, err
	}
	if roommates == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &roommates, err
}
