package user

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (h *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*Dto)
	err := h.CreateUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating user: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateUser(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into users (id, full_name) values ($1, $2);", dto.Id.String(), dto.FullName)
	return err
}
