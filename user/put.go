package user

import (
	"fmt"
	"net/http"
	"time"
)

// swagger:route PUT /v1/user/{id} user updateUser
// update a user
//
// responses:
//	201: noContentResponse
//  404: userErrorResponse
//  422: userErrorValidation

// Update handles PUT requests to update users
func (h *Handler) HandlePutUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(KeyId{}).(string)
	dto := r.Context().Value(KeyBody{}).(*UserDto)
	dto.Id = id
	err := h.UpdateUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating user: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateUser(dto *UserDto) error {
	err := h.db.Run(
		`update users set 
			id=$1, 
			full_name=$2, 
			user_status=$3,
			last_updated=$4,
			phone_number=$5, 
			nickname=$6, 
			bio=$7, 
			user_type=$8, 
			send_notifications=$9, 
			moving_reason=$10, 
			has_roommates=$11, 
			has_security_deposit=$12, 
			is_smoker=$13, 
			has_prev_lawsuit=$14, 
			has_prev_eviction=$15, 
			can_credit_check=$16, 
			has_pets=$17, 
			additional_details_general=$18, 
			move_in_date=$19, 
			move_out_date=$20, 
			additional_details_lease=$21,
			walkthrough_complete=$22
		where id=$1;`,
		dto.Id,
		dto.FullName,
		dto.UserStatus,
		time.Now().Format(time.RFC3339Nano),
		dto.PhoneNumber,
		dto.Nickname,
		dto.Bio,
		dto.UserType,
		dto.SendNotifications,
		dto.MovingReason,
		dto.HasRoommates,
		dto.HasSecurityDeposit,
		dto.IsSmoker,
		dto.HasPreviousLawsuit,
		dto.HasPreviousEviction,
		dto.CanCreditCheck,
		dto.HasPets,
		dto.AdditionalDetailsGeneral,
		dto.MoveInDate,
		dto.MoveOutDate,
		dto.AdditionalDetailsLease,
		dto.WalkthroughComplete,
	)
	return err
}
