package pet

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

func ConfigurePetRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *authenticator.Authenticator, authorizer *authorizer.Authorizer) {
	petHandler := NewHandler(logger, database, authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/pet/{id:%s}", util.REGEX_UUID), petHandler.HandleGetPet)
	getRouter.Use(authenticator.AuthenticateHttp, petHandler.GetIdFromPath, petHandler.CheckPermissionsPetIdView)

	getRouter2 := router.Methods(http.MethodGet).Subrouter()
	getRouter2.HandleFunc("/v1/pet", petHandler.HandleGetPetByUserId)
	getRouter2.Use(authenticator.AuthenticateHttp, petHandler.GetUserIdFromQueryParameters, petHandler.CheckPermissionsView)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/pet", petHandler.HandlePostPet)
	postRouter.Use(authenticator.AuthenticateHttp, petHandler.GetBody, petHandler.CheckPermissionsBodyEdit)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/pet/{id:%s}", util.REGEX_UUID), petHandler.HandlePutPet)
	putRouter.Use(authenticator.AuthenticateHttp, petHandler.GetBody, petHandler.GetIdFromPath, petHandler.CheckPermissionsPetIdEdit)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/pet/{id:%s}", util.REGEX_UUID), petHandler.HandleDeletePet)
	deleteRouter.Use(authenticator.AuthenticateHttp, petHandler.GetIdFromPath, petHandler.CheckPermissionsPetIdEdit)
}
