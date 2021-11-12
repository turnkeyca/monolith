package pet

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/key"
)

// swagger:route POST /v1/pet pet createPet
// create a new pet
//
// responses:
//	204: noContentResponse
//  400: petErrorResponse
//  422: petErrorResponse
//  500: petErrorResponse

// Create handles POST requests to add new pets
func (h *Handler) HandlePostPet(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*PetDto)
	err := h.CreatePet(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating pet: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreatePet(dto *PetDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run(
		`insert into pet (
			id, 
			user_id, 
			pet_type,
			breed, 
			size_type, 
			created_on,
			last_updated
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6,
			$6
		);`,
		dto.Id,
		dto.UserId,
		dto.PetType,
		dto.Breed,
		dto.SizeType,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
