package pet

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route GET /v1/pet/{id} pet getPet
// return a pet
// responses:
//	200: petResponse
//	404: petErrorResponse
//  500: petErrorResponse

// HandleGetPet handles GET requests
func (h *Handler) HandleGetPet(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	pet, err := h.GetPet(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting pet by id: %s, %s", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = pet.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

// swagger:route GET /v1/pet pet getPetsByUserId
// return all pets for a user
// responses:
//	200: petsResponse
//	500: petErrorResponse

// HandleGetPetByUserId handles GET requests
func (h *Handler) HandleGetPetByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyUserId{}).(string)
	pets, err := h.GetPetByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting pet by user id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(pets, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetPet(id string) (*PetDto, error) {
	var pets []PetDto
	err := h.db.Select(&pets, `select * from pet where id = $1;`, id)
	if err != nil {
		return nil, err
	}
	if pets == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(pets) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &pets[0], err
}

func (h *Handler) GetPetByUserId(userId string) (*[]PetDto, error) {
	var pets []PetDto
	err := h.db.Select(&pets, `select * from pet where user_id = $1;`, userId)
	if err != nil {
		return nil, err
	}
	if pets == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &pets, err
}
