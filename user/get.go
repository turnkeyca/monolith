package user

import (
	"fmt"
	"net/http"
)

// swagger:route GET /api/user/{id} user getUser
// return a user
// responses:
//	200: userResponse
//	404: userErrorResponse

// HandleGetUser handles GET requests
func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	user, err := h.GetUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = user.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetUser(id string) (*UserDto, error) {
	result, err := NewUserDatabase(h.db).SelectUser(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &result[0], err
}
