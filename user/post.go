package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// swagger:route POST /api/user user createUser
// create a new user
//
// responses:
//	204: noContentResponse
//  422: userErrorValidation
//  500: userErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*UserDto)
	err := h.CreateUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating user: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateUser(dto *UserDto) error {
	dto.Id = uuid.New().String()
	err := h.db.Run(
		`insert into users (
			id, 
			full_name, 
			email, 
			password, 
			user_status, 
			created_on, 
			last_updated,
			phone_number, 
			nickname, 
			bio, 
			user_type, 
			send_notifications, 
			moving_reason, 
			has_roommates, 
			has_security_deposit, 
			is_smoker, 
			has_prev_lawsuit, 
			has_prev_eviction, 
			can_credit_check, 
			has_pets, 
			additional_details_general, 
			move_in_date, 
			move_out_date, 
			additional_details_lease
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
			$10, 
			$11, 
			$12, 
			$13, 
			$14, 
			$15, 
			$16, 
			$17, 
			$18, 
			$19, 
			$20, 
			$21, 
			$22, 
			$23
		);`,
		dto.Id,
		dto.FullName,
		dto.Email,
		dto.Password,
		dto.UserStatus,
		time.Now().Format(time.RFC3339Nano),
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
