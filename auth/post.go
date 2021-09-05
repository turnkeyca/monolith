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
	token, err := GenerateToken(dto.LoginId)
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
	err = h.db.Run(
		`insert into users (
			id, 
			login_id, 
			user_status, 
			created_on, 
			last_updated
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5
		);`,
		id,
		dto.LoginId,
		"active",
		time.Now().Format(time.RFC3339Nano),
		time.Now().Format(time.RFC3339Nano),
	)
	if err != nil {
		return "", err
	}
	return id, nil
}
