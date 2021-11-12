package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/authenticator"
	"github.com/turnkeyca/monolith/authorizer"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/util"
)

type Handler struct {
	logger     *log.Logger
	db         *db.Database
	authorizer *authorizer.Authorizer
}

func NewHandler(logger *log.Logger, db *db.Database, authorizer *authorizer.Authorizer) *Handler {
	return &Handler{
		logger:     logger,
		db:         db,
		authorizer: authorizer,
	}
}

func ConfigureUserRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *authenticator.Authenticator, authorizer *authorizer.Authorizer) {
	userHandler := NewHandler(logger, database, authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/user/{id:%s}", util.REGEX_UUID), userHandler.HandleGetUser)
	getRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath, userHandler.CheckPermissionsView)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/user/{id:%s}", util.REGEX_UUID), userHandler.HandlePutUser)
	putRouter.Use(authenticator.AuthenticateHttp, userHandler.GetBody, userHandler.GetIdFromPath, userHandler.CheckPermissionsEdit)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/user/{id:%s}", util.REGEX_UUID), userHandler.HandleDeleteUser)
	deleteRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath, userHandler.CheckPermissionsEdit)
}
