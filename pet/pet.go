package pet

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

func ConfigurePetRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator, authorizer *permission.Authorizer) {
	petHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/pet/{id:%s}", util.REGEX_UUID), petHandler.HandleGetPet)
	getRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, petHandler.GetIdFromPath)
	getRouter.HandleFunc("/v1/pet", petHandler.HandleGetPetByUserId)
	getRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, petHandler.GetUserIdFromQueryParameters)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/pet", petHandler.HandlePostPet)
	postRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, petHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/pet/{id:%s}", util.REGEX_UUID), petHandler.HandlePutPet)
	putRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, petHandler.GetBody, petHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/pet/{id:%s}", util.REGEX_UUID), petHandler.HandleDeletePet)
	deleteRouter.Use(authenticator.AuthenticateHttp, authorizer.AuthorizeHttp, petHandler.GetIdFromPath)
}
