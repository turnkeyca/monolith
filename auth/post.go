package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

// swagger:route POST /api/auth registerNewToken
// create a new auth
//
// responses:
//	200: userIdResponse
//  500: authErrorResponse

// Create handles POST requests to add new users
func (h *Handler) HandleRegisterToken(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*RegisterTokenDto)
	id, err := h.getOrCreateUserId(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error registering token: %#v", err), http.StatusInternalServerError)
		return
	}
	err = h.saveToken(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error registering token: %#v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserId{Id: id})
}

func (h *Handler) getOrCreateUserId(dto *RegisterTokenDto) (string, error) {
	if dto.IsNewUser {
		id, err := h.CreateUser(dto)
		if err != nil {
			return "", err
		}
		return id, nil
	}
	name, err := getNameFromToken(dto.TokenString)
	if err != nil {
		return "", err
	}
	var id string
	err = h.db.Select(&id, `select id from users where name = $1`, name)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *Handler) CreateUser(dto *RegisterTokenDto) (string, error) {
	id := uuid.New().String()
	name, err := getNameFromToken(dto.TokenString)
	if err != nil {
		return "", err
	}
	err = h.db.Run(
		`insert into users (
			id, 
			name, 
			user_status, 
			created_on, 
			last_updated
		) values (
			$1, 
			$3, 
			$4, 
			$5, 
			$6
		);`,
		id,
		name,
		"active",
		time.Now().Format(time.RFC3339Nano),
		time.Now().Format(time.RFC3339Nano),
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *Handler) saveToken(token *RegisterTokenDto) error {
	id := uuid.New().String()
	return h.db.Run(`insert into token (
			id, 
			token, 
			status, 
			created_on
		) values (
			$1, 
			$2, 
			$3, 
			$4
		)`,
		id,
		token.TokenString,
		"active",
		time.Now().Format(time.RFC3339Nano),
	)
}

func getNameFromToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, nil)
	if err != nil {
		return "", err
	}
	claims := token.Claims.(jwt.MapClaims)
	return claims["name"].(string), nil
}
