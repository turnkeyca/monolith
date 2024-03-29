package reference

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route PUT /v1/reference/{id} reference updateReference
// update a reference
//
// responses:
//	204: noContentResponse
//  400: referenceErrorResponse
//  403: referenceErrorResponse
//  404: referenceErrorResponse
//  422: referenceErrorResponse
//  500: referenceErrorResponse

// Update handles PUT requests to update references
func (h *Handler) HandlePutReference(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*ReferenceDto)
	dto.Id = r.Context().Value(key.KeyId{}).(string)
	err := h.UpdateReference(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating reference: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateReference(dto *ReferenceDto) error {
	err := h.db.Run(
		`update reference set 
			full_name=$2, 
			email=$3, 
			phone_number=$4, 
			relationship=$5, 
			additional_details=$6, 
			last_updated=$7
		where id=$1;`,
		dto.Id,
		dto.FullName,
		dto.Email,
		dto.PhoneNumber,
		dto.Relationship,
		dto.AdditionalDetails,
		time.Now().Format(time.RFC3339Nano),
	)
	return err
}
