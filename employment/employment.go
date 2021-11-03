package employment

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

func ConfigureEmploymentRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator, authorizer *permission.Authorizer) {
	employmentHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/employment/{id:%s}", util.REGEX_UUID), employmentHandler.HandleGetEmployment)
	getRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetIdFromPath, employmentHandler.CheckPermissions)
	getRouter.HandleFunc("/v1/employment", employmentHandler.HandleGetEmploymentByUserId)
	getRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetUserIdFromQueryParameters, employmentHandler.CheckPermissions)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/employment", employmentHandler.HandlePostEmployment)
	postRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetBody, employmentHandler.CheckPermissions)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/employment/{id:%s}", util.REGEX_UUID), employmentHandler.HandlePutEmployment)
	putRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetBody, employmentHandler.GetIdFromPath, employmentHandler.CheckPermissions)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/employment/{id:%s}", util.REGEX_UUID), employmentHandler.HandleDeleteEmployment)
	deleteRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetIdFromPath, employmentHandler.CheckPermissions)
}
