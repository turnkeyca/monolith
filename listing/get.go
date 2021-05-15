package listing

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route GET /api/listing/{id} listing getListing
// return a listing
// responses:
//	200: listingResponse
//	404: listingErrorResponse

// HandleGetListing handles GET requests
func (h *Handler) HandleGetListing(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	listing, err := h.GetListing(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting listing by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = listing.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetListing(id uuid.UUID) (*Dto, error) {
	result, err := NewListingDatabase(h.db).SelectListing(id)
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
