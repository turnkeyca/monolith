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

func ConfigurePermissionRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	permissionHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandleGetPermission)
	getRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath)
	getRouter.HandleFunc("/v1/permission", permissionHandler.HandleGetPermissionByUserId)
	getRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetUserIdFromQueryParameters)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/permission", permissionHandler.HandlePostPermission)
	postRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandlePutPermission)
	putRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetBody, permissionHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/permission/{id:%s}", util.REGEX_UUID), permissionHandler.HandleDeletePermission)
	deleteRouter.Use(authenticator.AuthenticateHttp, permissionHandler.GetIdFromPath)
}