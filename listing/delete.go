package listing

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route DELETE /api/listing/{id} listing deleteListing
// delete a listing
//
// responses:
//	201: noContentResponse
//  404: listingErrorResponse
//  500: listingErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteListing(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	err := h.DeleteListing(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting listing by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteListing(id uuid.UUID) error {
	err := h.db.Run("delete from listings where id = $1;", id.String())
	return err
}
