package employment

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/key"
)

// swagger:route POST /v1/employment employment createEmployment
// create a new employment
//
// responses:
//	204: noContentResponse
//  400: employmentErrorResponse
//  403: employmentErrorResponse
//  422: employmentErrorResponse
//  500: employmentErrorResponse

// Create handles POST requests to add new employments
func (h *Handler) HandlePostEmployment(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*EmploymentDto)
	err := h.CreateEmployment(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating employment: %s", err), http.StatusInternalServerError)
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
			rent_affordability,
			annual_salary, 
			created_on,
			last_updated
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7,
			$8,
			$9,
			$9
		);`,
		dto.Id,
		dto.UserId,
		dto.Employer,
		dto.Occupation,
		dto.Duration,
		dto.AdditionalDetails,
		dto.RentAffordability,
		dto.AnnualSalary,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
