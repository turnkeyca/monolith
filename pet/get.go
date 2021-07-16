package pet

import (
	"fmt"
	"net/http"
)

// swagger:route GET /api/pet/{id} pet getPet
// return a pet
// responses:
//	200: petResponse
//	404: petErrorResponse

// HandleGetPet handles GET requests
func (h *Handler) HandleGetPet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
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

// swagger:route GET /api/pet pet getPetsByUserId
// return all pets ofr a user
// responses:
//	200: petsResponse
//	404: petErrorResponse

// HandleGetPetByUserId handles GET requests
func (h *Handler) HandleGetPetByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyUserId{}).(string)
	pets, err := h.GetPetByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting pet by user id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(pets, w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetPet(id string) (*PetDto, error) {
	result, err := NewPetDatabase(h.db).SelectPet(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &result[0], err
}

func (h *Handler) GetPetByUserId(userId string) (*[]PetDto, error) {
	result, err := NewPetDatabase(h.db).SelectPetsByUserId(userId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &result, err
}
