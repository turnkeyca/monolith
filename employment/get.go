package employment

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route GET /api/employment/{id} employment getEmployment
// return a employment
// responses:
//	200: employmentResponse
//	404: employmentErrorResponse

// HandleGetEmployment handles GET requests
func (h *Handler) HandleGetEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	employment, err := h.GetEmployment(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting employment by id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = employment.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetEmployment(id uuid.UUID) (*Dto, error) {
	result, err := NewEmploymentDatabase(h.db).SelectEmployment(id)
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
