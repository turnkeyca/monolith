package reference

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/permission"
	"github.com/turnkeyca/monolith/util"
)

type Handler struct {
	logger *log.Logger
	db     *db.Database
}

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

func ConfigureReferenceRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator, authorizer *permission.Authorizer) {
	referenceHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/reference/{id:%s}", util.REGEX_UUID), referenceHandler.HandleGetReference)
	getRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, referenceHandler.GetIdFromPath)
	getRouter.HandleFunc("/v1/reference", referenceHandler.HandleGetReferenceByUserId)
	getRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, referenceHandler.GetUserIdFromQueryParameters)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/reference", referenceHandler.HandlePostReference)
	postRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, referenceHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/reference/{id:%s}", util.REGEX_UUID), referenceHandler.HandlePutReference)
	putRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, referenceHandler.GetBody, referenceHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/reference/{id:%s}", util.REGEX_UUID), referenceHandler.HandleDeleteReference)
	deleteRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, referenceHandler.GetIdFromPath)
}
