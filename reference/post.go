package reference

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/turnkeyca/monolith/key"
)

// swagger:route POST /v1/reference reference createReference
// create a new reference
//
// responses:
//	204: noContentResponse
//  400: referenceErrorResponse
//  422: referenceErrorResponse
//  500: referenceErrorResponse

// Create handles POST requests to add new references
func (h *Handler) HandlePostReference(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*ReferenceDto)
	err := h.CreateReference(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating reference: %s", err), http.StatusInternalServerError)
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
			additional_details,
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
			$8
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
