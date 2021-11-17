package reference

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route GET /v1/reference/{id} reference getReference
// return a reference
// responses:
//	200: referenceResponse
//  403: referenceErrorResponse
//	404: referenceErrorResponse
//  500: referenceErrorResponse

// HandleGetReference handles GET requests
func (h *Handler) HandleGetReference(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	reference, err := h.GetReference(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting reference by id: %s, %s", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = reference.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

// swagger:route GET /v1/reference reference getReferencesByUserId
// return all references for a user
// responses:
//	200: referencesResponse
//  403: referenceErrorResponse
//	500: referenceErrorResponse

// HandleGetReferenceByUserId handles GET requests
func (h *Handler) HandleGetReferenceByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyUserId{}).(string)
	references, err := h.GetReferenceByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting reference by user id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(references, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetReference(id string) (*ReferenceDto, error) {
	var references []ReferenceDto
	err := h.db.Select(&references, `select * from reference where id = $1;`, id)
	if err != nil {
		return nil, err
	}
	if references == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(references) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &references[0], err
}

func (h *Handler) GetReferenceByUserId(userId string) (*[]ReferenceDto, error) {
	var references []ReferenceDto
	err := h.db.Select(&references, `select * from reference where user_id = $1;`, userId)
	if err != nil {
		return nil, err
	}
	if references == nil {
		return &[]ReferenceDto{}, nil
	}
	return &references, err
}
