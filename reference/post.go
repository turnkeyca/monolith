package reference

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// swagger:route POST /api/reference reference createReference
// create a new reference
//
// responses:
//	204: noContentResponse
//  422: referenceErrorValidation
//  500: referenceErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostReference(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*ReferenceDto)
	err := h.CreateReference(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating reference: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateReference(dto *ReferenceDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run(
		`insert into reference (
			id, 
			user_id, 
			full_name, 
			email, 
			phone_number, 
			relationship, 
			additional_details
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6, 
			$7
		);`,
		dto.Id,
		dto.UserId,
		dto.FullName,
		dto.Email,
		dto.PhoneNumber,
		dto.Relationship,
		dto.AdditionalDetails,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
