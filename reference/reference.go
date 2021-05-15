package reference

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

func ConfigureReferenceRoutes(regexUuid string, router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	referenceHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/reference/{id:%s}", regexUuid), referenceHandler.HandleGetReference)
	getRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetIdFromPath)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/reference", referenceHandler.HandlePostReference)
	postRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/reference/{id:%s}", regexUuid), referenceHandler.HandlePutReference)
	putRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetBody, referenceHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/reference/{id:%s}", regexUuid), referenceHandler.HandleDeleteReference)
	deleteRouter.Use(authenticator.AuthenticateHttp, referenceHandler.GetIdFromPath)
}
