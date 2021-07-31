package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// swagger:route POST /api/auth auth signUp
// create a new auth
//
// responses:
//	200: userIdResponse
//  500: authErrorResponse

// Create handles POST requests to add new users
func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*AuthDto)
	id, err := h.CreateUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating auth: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(UserId{Id: id})
}

// swagger:route POST /api/auth auth signIn
// create a new auth
//
// responses:
//	200: tokenResponse
//  500: authErrorResponse

// Create handles POST requests to add new users
func (h *Handler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	dto := r.Context().Value(KeyBody{}).(*AuthDto)
	token, err := h.VerifyUser(dto)
	if err != nil {
		http.Error(w, fmt.Sprintf("error verifying auth: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(*token)
}

func (h *Handler) CreateUser(dto *AuthDto) (string, error) {
	id := uuid.New().String()
	password, err := HashPassword(dto.Password)
	if err != nil {
		return "", err
	}
	err = h.db.Run(
		`insert into users (
			id, 
			email, 
			password, 
			user_status, 
			created_on, 
			last_updated
		) values (
			$1, 
			$2, 
			$3, 
			$4, 
			$5, 
			$6
		);`,
		id,
		dto.Email,
		password,
		"active",
		time.Now().Format(time.RFC3339Nano),
		time.Now().Format(time.RFC3339Nano),
	)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (h *Handler) VerifyUser(dto *AuthDto) (*Token, error) {
	var pwdRaw []string
	err := h.db.Select(&pwdRaw, "select password from users where email=$1", dto.Email)
	if err != nil {
		return nil, err
	}
	if len(pwdRaw) != 1 {
		return nil, fmt.Errorf("invalid result for user %s: unable to count users with these credentials", dto.Email)
	}
	pwd := pwdRaw[0]
	err = bcrypt.CompareHashAndPassword([]byte(pwd), []byte(dto.Password))
	if err != nil {
		return nil, err
	}
	t, err := GenerateJWT(dto.Email)
	if err != nil {
		return nil, err
	}
	return &Token{
		Email:       dto.Email,
		TokenString: t,
	}, nil
}
