package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

// swagger:route POST /v1/auth/registertoken auth registerNewToken
// register token
//
// responses:
//	200: tokenResponse
//  500: authErrorResponse

// Create handles POST requests to add new users
func (h *Handler) HandleRegisterToken(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*RegisterTokenDto)

	if dto.Secret != os.Getenv("SECRET_KEY") {
		http.Error(w, "error registering token: invalid secret key", http.StatusInternalServerError)
		return
	}
	id, err := h.getOrCreateUserId(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error registering token: %#v", err), http.StatusInternalServerError)
		return
	}
	token, err := GenerateToken(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error registering token: %#v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Token{Id: id, Token: token})
}

func (h *Handler) getOrCreateUserId(dto *RegisterTokenDto) (string, error) {
	if dto.IsNewUser {
		id, err := h.createUser(dto)
		if err != nil {
			return "", err
		}
		return id, nil
	}
	var id []string
	err := h.db.Select(&id, `select id from users where login_id = $1`, dto.LoginId)
	if len(id) != 1 {
		return "", fmt.Errorf("duplicate or nonexistent user")
	}
	if err != nil {
		return "", err
	}
	return id[0], nil
}

func (h *Handler) createUser(dto *RegisterTokenDto) (string, error) {
	var count []int
	err := h.db.Select(&count, `select count(*) from users where login_id = $1`, dto.LoginId)
	if err != nil {
		return "", err
	}
	if count[0] > 0 {
		return "", fmt.Errorf("duplicate users")
	}
	id := uuid.New().String()
	tx, err := h.db.Begin()
	if err != nil {
		return "", err
	}
	stmt, err := tx.Prepare(`insert into users (
		id, 
		login_id, 
		user_status, 
		created_on, 
		last_updated,
		full_name,
		email,
		phone_number,
		nickname,
		bio,
		user_type,
		moving_reason,
		additional_details_general,
		move_in_date,
		move_out_date,
		additional_details_lease,
		send_notifications,
		has_roommates,
		has_security_deposit,
		is_smoker,
		has_prev_lawsuit,
		has_prev_eviction,
		can_credit_check,
		has_pets,
		walkthrough_complete,
		terms_accepted
	) values (
		$1, $2, $3, $4, $4, $5, $5, $5, $5, $5, $5, $5, $5, $5, $5, $5, $6, $6, $6, $6, $6, $6, $6, $6, $6, $6
	)`)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		id,
		dto.LoginId,
		"active",
		time.Now().Format(time.RFC3339Nano),
		"",
		false,
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	stmt2, err := tx.Prepare(
		`insert into permission (
			id,
			user_id,
			on_user_id,
			permission,
			created_on,
			last_updated
		) values (
			$1, $2, $2, $3, $4, $4
		)`,
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt2.Close()
	_, err = stmt2.Exec(
		uuid.New().String(),
		id,
		"view",
		time.Now().Format(time.RFC3339Nano),
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	stmt3, err := tx.Prepare(
		`insert into permission (
			id,
			user_id,
			on_user_id,
			permission,
			created_on,
			last_updated
		) values (
			$1, $2, $2, $3, $4, $4
		)`,
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	defer stmt3.Close()
	_, err = stmt3.Exec(
		uuid.New().String(),
		id,
		"edit",
		time.Now().Format(time.RFC3339Nano),
	)
	if err != nil {
		tx.Rollback()
		return "", err
	}
	tx.Commit()
	return id, nil
}
