package reference

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route POST /api/reference reference createReference
// create a new reference
//
// responses:
//	200: referenceResponse
//  422: referenceErrorValidation
//  500: referenceErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostReference(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*Dto)
	err := h.CreateReference(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating reference: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateReference(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into references (id, full_name) values ($1, $2);", dto.Id.String(), dto.FullName)
	return err
}
