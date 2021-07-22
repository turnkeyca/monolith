package employment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// swagger:route POST /api/employment employment createEmployment
// create a new employment
//
// responses:
//	204: noContentResponse
//  422: employmentErrorValidation
//  500: employmentErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostEmployment(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*EmploymentDto)
	err := h.CreateEmployment(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating employment: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateEmployment(dto *EmploymentDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run(
		`insert into employment (
			id, 
			user_id, 
			employer, 
			occupation, 
			duration, 
			additional_details, 
			annual_salary, 
			created_on
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7,
			$8
		);`,
		dto.Id,
		dto.UserId,
		dto.Employer,
		dto.Occupation,
		dto.Duration,
		dto.AdditionalDetails,
		dto.AnnualSalary,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
