package listing

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route PUT /api/listing listing updateListing
// update a listing
//
// responses:
//	201: noContentResponse
//  404: listingErrorResponse
//  422: listingErrorValidation

// Update handles PUT requests to update listings
func (h *Handler) HandlePutListing(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	dto := r.Context().Value(KeyBody{}).(*Dto)
	dto.Id = id
	err := h.UpdateListing(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating listing: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateListing(dto *Dto) error {
	err := h.db.Run("update listings set id=$1, full_name=$2 where id=$1;", dto.Id.String())
	return err
}
