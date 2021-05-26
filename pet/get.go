package pet

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route GET /api/pet/{id} pet getPet
// return a pet
// responses:
//	200: petResponse
//	404: petErrorResponse

// HandleGetPet handles GET requests
func (h *Handler) HandleGetPet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	pet, err := h.GetPet(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting pet by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = pet.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetPet(id uuid.UUID) (*PetDto, error) {
	result, err := NewPetDatabase(h.db).SelectPet(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for id: %s", id.String())
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id.String())
	}
	return &result[0], err
}
