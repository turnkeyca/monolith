package authenticator

import (
	"context"
	"fmt"
	"net/http"

	"github.com/turnkeyca/monolith/key"
)

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error reading auth: %s", err), http.StatusBadRequest)
			return
		}
		ctx := context.WithValue(r.Context(), key.KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (a *Authenticator) AuthenticateHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			http.Error(rw, "invalid token: no token", http.StatusUnauthorized)
			return
		}
		claims, err := ValidateToken(r.Header["Token"][0])
		if err != nil {
			http.Error(rw, fmt.Sprintf("invalid token: %s", err), http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), key.KeyLoggedInUserId{}, claims.Id)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)
	})
}
