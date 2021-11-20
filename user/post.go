package user

import (
	"fmt"
	"net/http"
	"time"

	"github.com/turnkeyca/monolith/key"
)

// swagger:route POST /v1/user/{id}/activate user activateUser
// activate a user
//
// responses:
//	204: noContentResponse
//  403: userErrorResponse
//  500: userErrorResponse

// Activate handles POST requests to activate users
func (h *Handler) HandleActivateUser(w http.ResponseWriter, r *http.Request) {
	err := h.ActivateUser(r.Context().Value(key.KeyId{}).(string))
	if err != nil {
		http.Error(w, fmt.Sprintf("error activating user: %s", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) ActivateUser(id string) error {
	return h.db.Run(`update users set user_status=$3, last_updated=$2 where id=$1`, id, time.Now().Format(time.RFC3339Nano), ACTIVE)
}
