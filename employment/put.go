package employment

import (
	"fmt"
	"net/http"
	"time"
)

// swagger:route PUT /api/employment/{id} employment updateEmployment
// update a employment
//
// responses:
//	201: noContentResponse
//  404: employmentErrorResponse
//  422: employmentErrorValidation

// Update handles PUT requests to update employments
func (h *Handler) HandlePutEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	dto := r.Context().Value(KeyBody{}).(*EmploymentDto)
	dto.Id = id
	err := h.UpdateEmployment(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating employment: %#v\n", err), http.StatusInternalServerError)
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
