package employment

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/db"
)

type Handler struct {
	logger *log.Logger
	db     *db.Database
}

type KeyId struct{}
type KeyBody struct{}

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

func ConfigureEmploymentRoutes(regexUuid string, router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	employmentHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/employment/{id:%s}", regexUuid), employmentHandler.HandleGetEmployment)
	getRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetIdFromPath)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/employment", employmentHandler.HandlePostEmployment)
	postRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/employment/{id:%s}", regexUuid), employmentHandler.HandlePutEmployment)
	putRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetBody, employmentHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/employment/{id:%s}", regexUuid), employmentHandler.HandleDeleteEmployment)
	deleteRouter.Use(authenticator.AuthenticateHttp, employmentHandler.GetIdFromPath)
}
