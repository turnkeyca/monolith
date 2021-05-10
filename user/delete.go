package user

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route DELETE /api/user/{id} user deleteUser
// delete a user
//
// responses:
//	201: noContentResponse
//  404: userErrorResponse
//  500: userErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	err := h.DeleteUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting user by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteUser(id uuid.UUID) error {
	err := h.db.Run("delete from users where id = $1;", id.String())
	return err
}
