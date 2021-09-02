package auth

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type KeyBody struct{}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Println("here it goes")
		d, err := read(r.Body)
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

		_, err := ValidateToken(r.Header["Token"][0])
		if err != nil {
			http.Error(rw, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(rw, r)
	})
}

func read(r io.Reader) (*RegisterTokenDto, error) {
	d := RegisterTokenDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}
