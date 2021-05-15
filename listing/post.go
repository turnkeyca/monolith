package listing

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route POST /api/listing listing createListing
// create a new listing
//
// responses:
//	200: listingResponse
//  422: listingErrorValidation
//  500: listingErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostListing(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*Dto)
	err := h.CreateListing(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating listing: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateListing(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into listings (id, full_name) values ($1, $2);", dto.Id.String())
	return err
}
