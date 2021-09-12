package employment

import (
	"fmt"
	"net/http"
)

// swagger:route DELETE /v1/employment/{id} employment deleteEmployment
// delete a employment
//
// responses:
//	201: noContentResponse
//  404: employmentErrorResponse
//  500: employmentErrorResponse

// Delete handles DELETE requests and removes items from the database
func (h *Handler) HandleDeleteEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	err := h.DeleteEmployment(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error deleting employment by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) DeleteEmployment(id string) error {
	err := h.db.Run("delete from employment where id = $1;", id)
	return err
}
