package reference

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
		id := uuid.MustParse(mux.Vars(r)["id"]).String()
		ctx := context.WithValue(r.Context(), KeyId{}, id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetUserIdFromQueryParameters(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
			h.logger.Printf("decoding error: %#v", err)
			http.Error(w, "Error reading employment", http.StatusBadRequest)
		}
		err = d.Validate()
		if err != nil {
			h.logger.Printf("validation error: %#v", err)
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
