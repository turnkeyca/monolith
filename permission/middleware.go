package permission

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
			http.Error(w, "Error getting permission: missing query parameter userId", http.StatusBadRequest)
			return
		}
		userId := uuid.MustParse(r.URL.Query().Get("userId")).String()
		ctx := context.WithValue(r.Context(), key.KeyUserId{}, userId)
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
			http.Error(w, fmt.Sprintf("Error validating permission: %s", err), http.StatusUnprocessableEntity)
			return
		}
		ctx := context.WithValue(r.Context(), key.KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsUserIdView(next http.Handler) http.Handler {
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
		body := r.Context().Value(key.KeyBody{}).(*PermissionRequestDto)
		loggedInUserId := r.Context().Value(key.KeyLoggedInUserId{}).(string)
		err := h.authorizer.CheckUserIdAndToken(body.UserId, loggedInUserId, authorizer.EDIT)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsPermissionIdView(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.checkPermissionsPermissionId(r.Context().Value(key.KeyId{}).(string), r.Context().Value(key.KeyLoggedInUserId{}).(string), authorizer.VIEW)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) CheckPermissionsPermissionIdEdit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := h.checkPermissionsPermissionId(r.Context().Value(key.KeyId{}).(string), r.Context().Value(key.KeyLoggedInUserId{}).(string), authorizer.EDIT)
		if err != nil {
			http.Error(w, fmt.Sprintf("User does not have permission: %s", err), http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) checkPermissionsPermissionId(id string, loggedInUserId string, perm authorizer.PermissionType) error {
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
	if len(userId) <= 0 && len(onUserId) <= 0 {
		err = fmt.Errorf("permission not found")
	} else if len(userId) <= 0 && len(onUserId) > 0 {
		err = h.authorizer.CheckUserIdAndToken(onUserId[0], loggedInUserId, perm)
	} else if perm == authorizer.VIEW && len(userId) > 0 && len(onUserId) <= 0 {
		err = h.authorizer.CheckUserIdAndToken(userId[0], loggedInUserId, perm)
	} else {
		err = h.authorizer.CheckUserIdsAndTokenAny([]string{userId[0], onUserId[0]}, loggedInUserId, perm)
	}
	if err != nil {
		return fmt.Errorf("user does not have permission: %s", err)
	}
	return nil
}
