package employment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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
			http.Error(w, "Error reading employment", http.StatusBadRequest)
			return
		}
		err = d.Validate()
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Error validating employment: %s", err),
				http.StatusUnprocessableEntity,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
