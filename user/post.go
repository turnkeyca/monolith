package user

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route POST /api/user user createUser
// create a new user
//
// responses:
//	200: userResponse
//  422: userErrorValidation
//  500: userErrorResponse

// Create handles POST requests to add new products
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
