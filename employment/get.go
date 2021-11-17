package employment

import (
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route GET /v1/employment/{id} employment getEmployment
// return an employment
// responses:
//	200: employmentResponse
//  403: employmentErrorResponse
//  404: employmentErrorResponse
//  500: employmentErrorResponse

// HandleGetEmployment handles GET requests
func (h *Handler) HandleGetEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	employment, err := h.GetEmployment(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting employment by id: %s, %s", id, err), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = employment.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

// swagger:route GET /v1/employment employment getEmploymentsByUserId
// return employments for a user
// responses:
//	200: employmentsResponse
//  403: employmentErrorResponse
//	500: employmentErrorResponse

// HandleGetEmploymentByUserId handles GET requests
func (h *Handler) HandleGetEmploymentByUserId(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyUserId{}).(string)
	employments, err := h.GetEmploymentByUserId(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting employment by user id: %s, %s", id, err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = WriteAll(employments, w)
	if err != nil {
		http.Error(w, fmt.Sprintf("encoding error: %s", err), http.StatusInternalServerError)
	}
}

func (h *Handler) GetEmployment(id string) (*EmploymentDto, error) {
	var employments []EmploymentDto
	err := h.db.Select(&employments, `select * from employment where id = $1;`, id)
	if err != nil {
		return nil, err
	}
	if employments == nil {
		return nil, fmt.Errorf("no results for id: %s", id)
	}
	if len(employments) != 1 {
		return nil, fmt.Errorf("duplicate results for id: %s", id)
	}
	return &employments[0], err
}

func (h *Handler) GetEmploymentByUserId(userId string) (*[]EmploymentDto, error) {
	var employments []EmploymentDto
	err := h.db.Select(&employments, `select * from employment where user_id = $1;`, userId)
	if err != nil {
		return nil, err
	}
	if employments == nil {
		return &[]EmploymentDto{}, nil
	}
	return &employments, nil
}
