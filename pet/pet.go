package pet

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

func ConfigurePetRoutes(regexUuid string, router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	petHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/pet/{id:%s}", regexUuid), petHandler.HandleGetPet)
	getRouter.Use(authenticator.AuthenticateHttp, petHandler.GetIdFromPath)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/pet", petHandler.HandlePostPet)
	postRouter.Use(authenticator.AuthenticateHttp, petHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/pet/{id:%s}", regexUuid), petHandler.HandlePutPet)
	putRouter.Use(authenticator.AuthenticateHttp, petHandler.GetBody, petHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/pet/{id:%s}", regexUuid), petHandler.HandleDeletePet)
	deleteRouter.Use(authenticator.AuthenticateHttp, petHandler.GetIdFromPath)
}
