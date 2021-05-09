package user

import (
	"fmt"
	"net/http"
)

func (h *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(Dto)
	err := h.CreateUser(&dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating user: %#v\n", err), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateUser(dto *Dto) error {
	err := h.db.Run("insert into users (id) values ($1);", dto.Id.String())
	return err
}
