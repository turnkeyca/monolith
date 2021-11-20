package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route PUT /v1/user/{id} user updateUser
// update a user
//
// responses:
//	204: noContentResponse
//  400: userErrorResponse
//  403: userErrorResponse
//  422: userErrorResponse
//  500: userErrorResponse

// Update handles PUT requests to update users
func (h *Handler) HandlePutUser(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(key.KeyBody{}).(*UserDto)
	dto.Id = r.Context().Value(key.KeyId{}).(string)
	err := h.UpdateUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error updating user: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateUser(dto *UserDto) error {
	err := h.db.Run(
		`update users set 
			full_name=$2, 
			last_updated=$3,
			phone_number=$4, 
			nickname=$5, 
			bio=$6, 
			user_type=$7, 
			send_notifications=$8, 
			moving_reason=$9, 
			has_roommates=$10, 
			has_security_deposit=$11, 
			is_smoker=$12, 
			has_prev_lawsuit=$13, 
			has_prev_eviction=$14, 
			can_credit_check=$15, 
			has_pets=$16, 
			additional_details_general=$17,
			move_in_date=$18, 
			move_out_date=$19, 
			additional_details_lease=$20,
			walkthrough_complete=$21,
			terms_accepted=$22,
			email=$23
		where id=$1;`,
		dto.Id,
		dto.FullName,
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
		dto.AcceptedTerms,
		dto.Email,
	)
	return err
}
