package employment

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route POST /api/employment employment createEmployment
// create a new employment
//
// responses:
//	200: employmentResponse
//  422: employmentErrorValidation
//  500: employmentErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostEmployment(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*Dto)
	err := h.CreateEmployment(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating employment: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateEmployment(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into employments (id, full_name) values ($1, $2);", dto.Id.String(), dto.FullName)
	return err
}
