package reference

import (
	"fmt"
	"net/http"
)

// swagger:route GET /api/reference/{id} reference getReference
// return a reference
// responses:
//	200: referenceResponse
//	404: referenceErrorResponse

// HandleGetReference handles GET requests
func (h *Handler) HandleGetReference(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	reference, err := h.GetReference(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting reference by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = reference.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

// swagger:route GET /api/reference reference getReferencesByUserId
// return all references ofr a user
// responses:
//	200: referencesResponse
//	404: referenceErrorResponse

// HandleGetReferenceByUserId handles GET requests
func (h *Handler) HandleGetReferenceByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyUserId{}).(string)
	references, err := h.GetReferenceByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting reference by user id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(references, w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetReference(id string) (*ReferenceDto, error) {
	result, err := NewReferenceDatabase(h.db).SelectReference(id)
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

func (h *Handler) GetReferenceByUserId(userId string) (*[]ReferenceDto, error) {
	result, err := NewReferenceDatabase(h.db).SelectReferencesByUserId(userId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &result, err
}
