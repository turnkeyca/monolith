package reference

import (
	"fmt"
	"net/http"
	"time"
)

// swagger:route PUT /api/reference/{id} reference updateReference
// update a reference
//
// responses:
//	201: noContentResponse
//  404: referenceErrorResponse
//  422: referenceErrorValidation

// Update handles PUT requests to update references
func (h *Handler) HandlePutReference(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	dto := r.Context().Value(KeyBody{}).(*ReferenceDto)
	dto.Id = id
	err := h.UpdateReference(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating reference: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateReference(dto *ReferenceDto) error {
	err := h.db.Run(
		`update reference set 
			id=$1, 
			user_id=$2, 
			full_name=$3, 
			email=$4, 
			phone_number=$5, 
			relationship=$6, 
			additional_details=$7, 
			last_updated=$8
		where id=$1;`,
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
