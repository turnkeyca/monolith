package user

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// swagger:route POST /api/user user createUser
// create a new user
//
// responses:
//	200: userResponse
//  422: userErrorValidation
//  500: userErrorResponse

// Create handles POST requests to add new products
func (h *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*Dto)
	err := h.CreateUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating user: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) CreateUser(dto *Dto) error {
	dto.Id = uuid.New()
	err := h.db.Run("insert into users (id, full_name, email, password, phone_number, nickname, bio, city, province, user_type, send_notifications, moving_reason, has_roommates, has_security_deposit, is_smoker, has_prev_lawsuit, has_prev_eviction, can_credit_check, has_pets, additional_details, move_in_date, move_out_date, property_management_company, additional_details_lease, monthly_budget_min, monthly_budget_max) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26);", dto.Id, dto.FullName, dto.Email, dto.Password, dto.PhoneNumber, dto.Nickname, dto.Bio, dto.City, dto.Province, dto.UserType, dto.SendNotifications, dto.MovingReason, dto.HasRoommates, dto.HasSecurityDeposit, dto.IsSmoker, dto.HasPreviousLawsuit, dto.HasPreviousEviction, dto.CanCreditCheck, dto.HasPets, dto.AdditionalDetails, dto.MoveInDate, dto.MoveOutDate, dto.PropertyManagementCompany, dto.AdditionalDetailsLease, dto.MonthlyBudgetMin, dto.MonthlyBudgetMax)
	return err
}
