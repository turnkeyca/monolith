package roommate

import (
	"fmt"
	"net/http"
)

// swagger:route PUT /api/roommate roommate updateRoommate
// update a roommate
//
// responses:
//	201: noContentResponse
//  404: roommateErrorResponse
//  422: roommateErrorValidation

// Update handles PUT requests to update roommates
func (h *Handler) HandlePutRoommate(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	dto := r.Context().Value(KeyBody{}).(*Dto)
	dto.Id = id
	err := h.UpdateRoommate(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating roommate: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateRoommate(dto *Dto) error {
	err := h.db.Run("update roommate set id=$1, user_id=$2, full_name=$3, email=$4, additional_details=$5 where id=$1;", dto.Id, dto.UserId, dto.FullName, dto.Email, dto.AdditionalDetails)
	return err
}
