package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/util"
)

type Handler struct {
	logger *log.Logger
	db     *db.Database
}

type KeyId struct{}
type KeyBody struct{}

func NewHandler(logger *log.Logger, db *db.Database) *Handler {
	return &Handler{
		logger: logger,
		db:     db,
	}
}

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func ConfigureUserRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	userHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", util.REGEX_UUID), userHandler.HandleGetUser)
	getRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", util.REGEX_UUID), userHandler.HandlePutUser)
	putRouter.Use(authenticator.AuthenticateHttp, userHandler.GetBody, userHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", util.REGEX_UUID), userHandler.HandleDeleteUser)
	deleteRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath)
}
