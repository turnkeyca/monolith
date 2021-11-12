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
	logger     *log.Logger
	authorizer *permission.Authorizer
	db         *db.Database
}

func NewHandler(logger *log.Logger, db *db.Database, authorizer *permission.Authorizer) *Handler {
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

func ConfigureEmploymentRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator, authorizer *permission.Authorizer) {
	employmentHandler := NewHandler(logger, database, authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/employment/{id:%s}", util.REGEX_UUID), employmentHandler.HandleGetEmployment)
	getRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetIdFromPath, employmentHandler.CheckPermissionsEmploymentIdView)
	getRouter2 := router.Methods(http.MethodGet).Subrouter()
	getRouter2.HandleFunc("/v1/employment", employmentHandler.HandleGetEmploymentByUserId)
	getRouter2.Use(authenticator.AuthenticateHttp, employmentHandler.GetUserIdFromQueryParameters, employmentHandler.CheckPermissionsView)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/employment", employmentHandler.HandlePostEmployment)
	postRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetBody, employmentHandler.CheckPermissionsBodyEdit)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/employment/{id:%s}", util.REGEX_UUID), employmentHandler.HandlePutEmployment)
	putRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetBody, employmentHandler.GetIdFromPath, employmentHandler.CheckPermissionsEmploymentIdEdit)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/employment/{id:%s}", util.REGEX_UUID), employmentHandler.HandleDeleteEmployment)
	deleteRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetIdFromPath, employmentHandler.CheckPermissionsEmploymentIdEdit)
}
