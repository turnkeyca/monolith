package permission

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/util"
)

type Authorizer struct {
	logger *log.Logger
	db     *db.Database
}

type Handler struct {
	logger     *log.Logger
	authorizer *Authorizer
	db         *PermissionDatabase
}

func New(logger *log.Logger, db *db.Database) *Authorizer {
	return &Authorizer{
		logger: logger,
		db:     db,
	}
}

func NewHandler(logger *log.Logger, db *PermissionDatabase, authorizer *Authorizer) *Handler {
	return &Handler{
		logger:     logger,
		db:         db,
		authorizer: authorizer,
	}
}

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func ConfigurePermissionRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator, authorizer *Authorizer) {
	permissionHandler := NewHandler(logger, NewPermissionDatabase(database), authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandleGetPermission)
	getRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath, permissionHandler.CheckPermissionsWithPermissionIdView)
	getRouter2 := router.Methods(http.MethodGet).Subrouter()
	getRouter2.HandleFunc("/v1/permission", permissionHandler.HandleGetPermissionByUserId)
	getRouter2.Use(authenticator.AuthenticateHttp, permissionHandler.GetUserIdFromQueryParameters, permissionHandler.CheckPermissionsView)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/permission", permissionHandler.HandlePostPermission)
	postRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetRequestBody)

	postRouter2 := router.Methods(http.MethodPost).Subrouter()
	postRouter2.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}/accept", util.REGEX_UUID), permissionHandler.HandleAcceptPermission)
	postRouter2.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandleDeletePermission)
	deleteRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath, permissionHandler.CheckPermissionsWithPermissionIdEdit)
}
