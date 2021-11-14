package reference

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
	authorizer *authorizer.Authorizer
	db         *db.Database
}

func NewHandler(logger *log.Logger, db *db.Database, authorizer *authorizer.Authorizer) *Handler {
	return &Handler{
		logger:     logger,
		db:         db,
		authorizer: authorizer,
	}
}

func ConfigureReferenceRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *authenticator.Authenticator, authorizer *authorizer.Authorizer) {
	referenceHandler := NewHandler(logger, database, authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/reference/{id:%s}", util.REGEX_UUID), referenceHandler.HandleGetReference)
	getRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetIdFromPath, referenceHandler.CheckPermissionsReferenceIdView)

	getRouter2 := router.Methods(http.MethodGet).Subrouter()
	getRouter2.HandleFunc("/v1/reference", referenceHandler.HandleGetReferenceByUserId)
	getRouter2.Use(authenticator.AuthenticateHttp, referenceHandler.GetUserIdFromQueryParameters, referenceHandler.CheckPermissionsView)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/reference", referenceHandler.HandlePostReference)
	postRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetBody, referenceHandler.CheckPermissionsBodyEdit)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/reference/{id:%s}", util.REGEX_UUID), referenceHandler.HandlePutReference)
	putRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetBody, referenceHandler.GetIdFromPath, referenceHandler.CheckPermissionsReferenceIdEdit)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/reference/{id:%s}", util.REGEX_UUID), referenceHandler.HandleDeleteReference)
	deleteRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetIdFromPath, referenceHandler.CheckPermissionsReferenceIdEdit)
}
