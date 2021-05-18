package employment

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

// swagger:route PUT /api/employment employment updateEmployment
// update a employment
//
// responses:
//	201: noContentResponse
//  404: employmentErrorResponse
//  422: employmentErrorValidation

// Update handles PUT requests to update employments
func (h *Handler) HandlePutEmployment(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	dto := r.Context().Value(KeyBody{}).(Dto)
	dto.Id = id
	err := h.UpdateEmployment(&dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating employment: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateEmployment(dto *Dto) error {
	err := h.db.Run("update employment set id=$1, user_id=$2, employer=$3, occupation=$4, duration=$5, additional_details=$6, annual_salary=$7 where id=$1;", dto.Id.String(), dto.UserId.String(), dto.Employer, dto.Occupation, dto.Duration, dto.AdditionalDetails, strconv.FormatFloat(dto.AnnualSalary, 'f', 2, 64))
	return err
}
