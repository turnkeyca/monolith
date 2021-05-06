package user

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	user, err := h.GetUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user by id: %s, %#v\n", id, err), http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = user.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetUser(id uuid.UUID) (*Dto, error) {
	result, err := h.db.Query("select * from users where id = %s;", id.String())
	return Assemble(result), err
}
