package authenticator

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/db"
)

type Authenticator struct {
	logger *log.Logger
	db     *db.Database
}

type Handler struct {
	logger *log.Logger
	db     *db.Database
}

func New(logger *log.Logger, db *db.Database) *Authenticator {
	return &Authenticator{
		logger: logger,
		db:     db,
	}
}

func NewHandler(logger *log.Logger, db *db.Database) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

func ConfigureAuthRoutes(router *mux.Router, logger *log.Logger, database *db.Database) {
	authHandler := NewHandler(logger, database)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/auth/registertoken", authHandler.HandleRegisterToken)
	postRouter.Use(authHandler.GetBody)
}
