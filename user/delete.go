package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route DELETE /v1/user/{id} user deleteUser
// delete a user
//
// responses:
//	204: noContentResponse
//  500: userErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.DeleteUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting user by id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteUser(id string) error {
	return h.db.Run(`update users set user_status='inactive', last_updated=$2 where id = $1;`, id, time.Now().Format(time.RFC3339Nano))
}
