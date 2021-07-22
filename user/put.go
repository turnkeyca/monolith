package user

import (
	"fmt"
	"net/http"
	"time"
)

// swagger:route PUT /api/user/{id} user updateUser
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
			email=$3, 
			password=$4, 
			user_status=$5
			last_updated=$6
			phone_number=$7, 
			nickname=$8, 
			bio=$9, 
			user_type=$10, 
			send_notifications=$11, 
			moving_reason=$12, 
			has_roommates=$13, 
			has_security_deposit=$14, 
			is_smoker=$15, 
			has_prev_lawsuit=$16, 
			has_prev_eviction=$17, 
			can_credit_check=$18, 
			has_pets=$19, 
			additional_details_general=$20, 
			move_in_date=$21, 
			move_out_date=$22, 
			additional_details_lease=$23, 
		where id=$1;`,
		dto.Id,
		dto.FullName,
		dto.Email,
		dto.Password,
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
	)
	return err
}
