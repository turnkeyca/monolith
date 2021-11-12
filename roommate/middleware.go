package roommate

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

func (h *Handler) GetUserIdFromQueryParameters(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("userId") == "" {
			http.Error(w, "Error getting roommates: missing query parameter userId", http.StatusBadRequest)
			return
		}
		userId := uuid.MustParse(r.URL.Query().Get("userId")).String()
		ctx := context.WithValue(r.Context(), key.KeyUserId{}, userId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading roommate: %s", err), http.StatusBadRequest)
			return
		}
		err = d.Validate()
		if err != nil {
			http.Error(w, fmt.Sprintf("Error validating roommate: %s", err), http.StatusUnprocessableEntity)
			return
		}
		ctx := context.WithValue(r.Context(), key.KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.Context().Value(key.KeyUserId{}).(string)
		loggedInUserId := r.Context().Value(key.KeyLoggedInUserId{}).(string)
		err := h.authorizer.CheckUserIdAndToken(id, loggedInUserId, authorizer.VIEW)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsBodyEdit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := r.Context().Value(key.KeyBody{}).(*RoommateDto)
		loggedInUserId := r.Context().Value(key.KeyLoggedInUserId{}).(string)
		err := h.authorizer.CheckUserIdAndToken(body.UserId, loggedInUserId, authorizer.VIEW)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsRoommateIdEdit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id []string
		err := h.db.Select(&id, `select user_id from roommate where id=$1;`, r.Context().Value(key.KeyId{}).(string))
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		loggedInUserId := r.Context().Value(key.KeyLoggedInUserId{}).(string)
		err = h.authorizer.CheckUserIdAndToken(id[0], loggedInUserId, authorizer.EDIT)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsRoommateIdView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var id []string
		err := h.db.Select(&id, `select user_id from roommate where id=$1;`, r.Context().Value(key.KeyId{}).(string))
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		loggedInUserId := r.Context().Value(key.KeyLoggedInUserId{}).(string)
		err = h.authorizer.CheckUserIdAndToken(id[0], loggedInUserId, authorizer.VIEW)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
