package employment

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route DELETE /api/employment/{id} employment deleteEmployment
// delete a employment
//
// responses:
//	201: noContentResponse
//  404: employmentErrorResponse
//  500: employmentErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	err := h.DeleteEmployment(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting employment by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteEmployment(id uuid.UUID) error {
	err := h.db.Run("delete from employments where id = $1;", id.String())
	return err
}
