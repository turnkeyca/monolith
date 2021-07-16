package employment

import (
	"fmt"
	"net/http"
)

// swagger:route GET /api/employment/{id} employment getEmployment
// return an employment
// responses:
//	200: employmentResponse
//	404: employmentErrorResponse

// HandleGetEmployment handles GET requests
func (h *Handler) HandleGetEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
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

// swagger:route GET /api/employment employment getEmploymentsByUserId
// return employments for a user
// responses:
//	200: employmentsResponse
//	404: employmentErrorResponse

// HandleGetEmploymentByUserId handles GET requests
func (h *Handler) HandleGetEmploymentByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyUserId{}).(string)
	employments, err := h.GetEmploymentByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting employment by user id: %s, %#v\n", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(employments, w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetEmployment(id string) (*EmploymentDto, error) {
	result, err := NewEmploymentDatabase(h.db).SelectEmployment(id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(result) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &result[0], err
}

func (h *Handler) GetEmploymentByUserId(userId string) (*[]EmploymentDto, error) {
	result, err := NewEmploymentDatabase(h.db).SelectEmploymentByUserId(userId)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, fmt.Errorf("no results for user id: %s", userId)
	}
	return &result, err
}
