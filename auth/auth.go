package auth

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/db"
)

type GenericError struct {
	Message string `json:"message"`
}

type Authenticator struct {
	logger *log.Logger
}

type Handler struct {
	logger *log.Logger
	db     *db.Database
}

func New(logger *log.Logger) *Authenticator {
	return &Authenticator{
		logger: logger,
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
	postRouter.HandleFunc("/api/auth/signup", authHandler.HandleSignUp)
	postRouter.HandleFunc("/api/auth/signin", authHandler.HandleSignIn)
	postRouter.Use(authHandler.GetBody)
}
