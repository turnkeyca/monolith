package pet

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route PUT /api/pet pet updatePet
// update a pet
//
// responses:
//	201: noContentResponse
//  404: petErrorResponse
//  422: petErrorValidation

// Update handles PUT requests to update pets
func (h *Handler) HandlePutPet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	dto := r.Context().Value(KeyBody{}).(*Dto)
	dto.Id = id
	err := h.UpdatePet(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating pet: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdatePet(dto *Dto) error {
	err := h.db.Run("update pet set id=$1, full_name=$2 where id=$1;", dto.Id.String())
	return err
}
