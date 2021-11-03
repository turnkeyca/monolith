package auth

import (
	"context"
	"net/http"
)

type KeyBody struct{}
type KeyLoggedInUserId struct{}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			http.Error(w, "error reading auth", http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (a *Authenticator) AuthenticateHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			http.Error(rw, "no token", http.StatusUnauthorized)
			return
		}

		claims, err := ValidateToken(r.Header["Token"][0])
		if err != nil {
			http.Error(rw, "invalid token", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), KeyLoggedInUserId{}, claims.LoginId)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
