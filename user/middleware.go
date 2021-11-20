package user

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/authorizer"
	"github.com/turnkeyca/monolith/key"
)

func (h *Handler) GetIdFromPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.MustParse(mux.Vars(r)["id"]).String()
		ctx := context.WithValue(r.Context(), key.KeyId{}, id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading user: %s", err), http.StatusBadRequest)
			return
		}
		err = d.Validate()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error validating user: %s", err), http.StatusUnprocessableEntity)
			return
		}
		ctx := context.WithValue(r.Context(), key.KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.authorizer.CheckUserIdAndToken(
			r.Context().Value(key.KeyId{}).(string),
			r.Context().Value(key.KeyLoggedInUserId{}).(string),
			authorizer.VIEW,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsEdit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.authorizer.CheckUserIdAndToken(
			r.Context().Value(key.KeyId{}).(string),
			r.Context().Value(key.KeyLoggedInUserId{}).(string),
			authorizer.EDIT,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsEditNoInactiveCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.authorizer.CheckUserIdAndTokenNoActiveCheck(
			r.Context().Value(key.KeyId{}).(string),
			r.Context().Value(key.KeyLoggedInUserId{}).(string),
			authorizer.EDIT,
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
