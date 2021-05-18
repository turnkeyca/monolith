package pet

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// swagger:route POST /api/pet pet createPet
// create a new pet
//
// responses:
//	200: petResponse
//  422: petErrorValidation
//  500: petErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostPet(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(Dto)
	err := h.CreatePet(&dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating pet: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreatePet(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into pet (id, user_id, breed, weight) values ($1, $2, $3, $4);", dto.Id.String(), dto.UserId.String(), dto.Breed, strconv.FormatFloat(dto.Weight, 'f', 2, 64))
	return err
}
