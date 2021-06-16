package auth

import (
	"log"
	"net/http"
)

type Authenticator struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Authenticator {
	return &Authenticator{
		logger: logger,
	}
}

func (a *Authenticator) AuthenticateHttp(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		a.logger.Println("authenticated")
		next.ServeHTTP(rw, r)
	})
}
