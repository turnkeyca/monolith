package permission

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

func ConfigurePermissionRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *authenticator.Authenticator, authorizer *authorizer.Authorizer) {
	permissionHandler := NewHandler(logger, database, authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandleGetPermission)
	getRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath, permissionHandler.CheckPermissionsPermissionIdView)

	getRouter2 := router.Methods(http.MethodGet).Subrouter()
	getRouter2.HandleFunc("/v1/permission", permissionHandler.HandleGetPermissionByUserId)
	getRouter2.Use(authenticator.AuthenticateHttp, permissionHandler.GetUserIdFromQueryParameters, permissionHandler.CheckPermissionsUserIdView)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/permission", permissionHandler.HandlePostPermission)
	postRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetRequestBody)

	postRouter2 := router.Methods(http.MethodPost).Subrouter()
	postRouter2.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}/accept", util.REGEX_UUID), permissionHandler.HandleAcceptPermission)
	postRouter2.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandleDeletePermission)
	deleteRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath, permissionHandler.CheckPermissionsPermissionIdEdit)
}
