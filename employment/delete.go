package employment

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route DELETE /v1/employment/{id} employment deleteEmployment
// delete a employment
//
// responses:
//	204: noContentResponse
//  500: employmentErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	err := h.DeleteEmployment(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting employment by id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteEmployment(id string) error {
	return h.db.Run(`delete from employment where id = $1;`, id)
}
