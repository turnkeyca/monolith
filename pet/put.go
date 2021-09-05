package pet

import (
	"fmt"
	"net/http"
	"time"
)

// swagger:route PUT /v1/pet/{id} pet updatePet
// update a pet
//
// responses:
//	201: noContentResponse
//  404: petErrorResponse
//  422: petErrorValidation

// Update handles PUT requests to update pets
func (h *Handler) HandlePutPet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	dto := r.Context().Value(KeyBody{}).(*PetDto)
	dto.Id = id
	err := h.UpdatePet(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating pet: %#v\n", err), http.StatusInternalServerError)
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
