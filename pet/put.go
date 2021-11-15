package pet

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route PUT /v1/pet/{id} pet updatePet
// update a pet
//
// responses:
//	204: noContentResponse
//  400: petErrorResponse
//  403: petErrorResponse
//  422: petErrorResponse
//  500: petErrorResponse

// Update handles PUT requests to update pets
func (h *Handler) HandlePutPet(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*PetDto)
	dto.Id = r.Context().Value(key.KeyId{}).(string)
	err := h.UpdatePet(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating pet: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdatePet(dto *PetDto) error {
	err := h.db.Run(
		`update pet set 
			id=$1, 
			user_id=$2, 
			pet_type=$3,
			breed=$4, 
			size_type=$5,
			last_updated=$6
		where id=$1;`,
		dto.Id,
		dto.UserId,
		dto.PetType,
		dto.Breed,
		dto.SizeType,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
