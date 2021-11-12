package permission

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
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

func (h *Handler) GetRequestBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := ReadPermissionRequestDto(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading permission: %s", err), http.StatusBadRequest)
			return
		}
		err = d.Validate()
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error validating permission: %s", err),
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
		err := h.authorizer.CheckUserIdAndToken(id, loggedInUserId, VIEW)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsWithPermissionIdView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.checkPermissionsWithPermissionId(r.Context().Value(KeyId{}).(string), r.Context().Value(auth.KeyLoggedInUserId{}).(string), VIEW)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsWithPermissionIdEdit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.checkPermissionsWithPermissionId(r.Context().Value(KeyId{}).(string), r.Context().Value(auth.KeyLoggedInUserId{}).(string), EDIT)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) checkPermissionsWithPermissionId(id string, loggedInUserId string, perm PermissionType) error {
	var userId []string
	err := h.db.Select(&userId, `select user_id from permission where id=$1;`, id)
	if err != nil {
		return fmt.Errorf("user does not have permission: %s", err)
	}
	var onUserId []string
	err = h.db.Select(&onUserId, `select on_user_id from permission where id=$1;`, id)
	if err != nil {
		return fmt.Errorf("user does not have permission: %s", err)
	}
	err = h.authorizer.CheckUserIdsAndTokenAny([]string{userId[0], onUserId[0]}, loggedInUserId, perm)
	if err != nil {
		return fmt.Errorf("user does not have permission: %s", err)
	}
	return nil
}
