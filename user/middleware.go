package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/permission"
)

func (h *Handler) GetIdFromPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.MustParse(mux.Vars(r)["id"]).String()
		ctx := context.WithValue(r.Context(), KeyId{}, id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			http.Error(w, "Error reading user", http.StatusBadRequest)
			return
		}
		err = d.Validate()
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error validating user: %s", err),
				http.StatusUnprocessableEntity,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissions(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(KeyId{}).(string)
		loggedInUserId := r.Context().Value(auth.KeyLoggedInUserId{}).(string)
		err := permission.CheckUserIdAndToken(id, loggedInUserId)
		if err != nil {
			http.Error(w, "User does not have permission", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
