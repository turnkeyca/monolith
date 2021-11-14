package employment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route PUT /v1/employment/{id} employment updateEmployment
// update a employment
//
// responses:
//	204: noContentResponse
//  400: employmentErrorResponse
//  404: employmentErrorResponse
//  422: employmentErrorResponse
//  500: employmentErrorResponse

// Update handles PUT requests to update employments
func (h *Handler) HandlePutEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(key.KeyId{}).(string)
	dto := r.Context().Value(key.KeyBody{}).(*EmploymentDto)
	dto.Id = id
	err := h.UpdateEmployment(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating employment: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateEmployment(dto *EmploymentDto) error {
	err := h.db.Run(
		`update employment set 
			id=$1, 
			user_id=$2, 
			employer=$3, 
			occupation=$4, 
			duration=$5, 
			additional_details=$6, 
			annual_salary=$7,
			rent_affordability=$8,
			last_updated=$9
		where id=$1;`,
		dto.Id,
		dto.UserId,
		dto.Employer,
		dto.Occupation,
		dto.Duration,
		dto.AdditionalDetails,
		dto.AnnualSalary,
		dto.RentAffordability,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
