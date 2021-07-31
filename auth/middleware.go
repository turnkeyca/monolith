package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type KeyBody struct{}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		var mySigningKey = []byte(os.Getenv("SECRET_KEY"))

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, fmt.Errorf("error parsing token")
			}
			return mySigningKey, nil
		})

		if err != nil {
			http.Error(rw, "expired token", http.StatusUnauthorized)
			return
		}
		if !token.Valid {
			http.Error(rw, "invalid token", http.StatusUnauthorized)
		}
		next.ServeHTTP(rw, r)
	})
}

func read(r io.Reader) (*AuthDto, error) {
	d := AuthDto{}
	err := json.NewDecoder(r).Decode(&d)
	return &d, err
}
