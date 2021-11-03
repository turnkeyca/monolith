package pet

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/permission"
)

type KeyId struct{}
type KeyBody struct{}
type KeyUserId struct{}

func (h *Handler) GetIdFromPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if mux.Vars(r)["id"] == "" {
			next.ServeHTTP(w, r)
			return
		}
		id := uuid.MustParse(mux.Vars(r)["id"]).String()
		ctx := context.WithValue(r.Context(), KeyId{}, id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetUserIdFromQueryParameters(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("userId") == "" {
			next.ServeHTTP(w, r)
			return
		}
		userId := uuid.MustParse(r.URL.Query().Get("userId")).String()
		ctx := context.WithValue(r.Context(), KeyUserId{}, userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			http.Error(w, "Error reading pet", http.StatusBadRequest)
			return
		}
		err = d.Validate()
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error validating pet: %s", err),
				http.StatusUnprocessableEntity,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(KeyUserId{}).(string)
		loggedInUserId := r.Context().Value(auth.KeyLoggedInUserId{}).(string)
		err := h.authorizer.CheckUserIdAndToken(id, loggedInUserId, permission.VIEW)
		if err != nil {
			http.Error(w, "User does not have permission", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsPetIdEdit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id string
		err := h.db.Select(&id, `select user_id from pet where id=$1;`, r.Context().Value(KeyId{}).(string))
		if err != nil {
			http.Error(w, "User does not have permission", http.StatusForbidden)
			return
		}
		loggedInUserId := r.Context().Value(auth.KeyLoggedInUserId{}).(string)
		err = h.authorizer.CheckUserIdAndToken(id, loggedInUserId, permission.EDIT)
		if err != nil {
			http.Error(w, "User does not have permission", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsPetIdView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id string
		err := h.db.Select(&id, `select user_id from pet where id=$1;`, r.Context().Value(KeyId{}).(string))
		if err != nil {
			http.Error(w, "User does not have permission", http.StatusForbidden)
			return
		}
		loggedInUserId := r.Context().Value(auth.KeyLoggedInUserId{}).(string)
		err = h.authorizer.CheckUserIdAndToken(id, loggedInUserId, permission.VIEW)
		if err != nil {
			http.Error(w, "User does not have permission", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
