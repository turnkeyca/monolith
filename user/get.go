package user

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route GET /v1/user/{id} user getUser
// return a user
// responses:
//	200: userResponse
//	404: userErrorResponse
//  500: userErrorResponse

// HandleGetUser handles GET requests
func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	user, err := h.GetUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user by id: %s, %s", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = user.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetUser(id string) (*UserDto, error) {
	var users []UserDto
	err := h.db.Select(&users, "select * from users where id = $1;", id)
	if err != nil {
		return nil, err
	}
	if users == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(users) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &users[0], err
}
