package listing

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

func ConfigureListingRoutes(regexUuid string, router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	listingHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/listing/{id:%s}", regexUuid), listingHandler.HandleGetListing)
	getRouter.Use(authenticator.AuthenticateHttp, listingHandler.GetIdFromPath)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/listing", listingHandler.HandlePostListing)
	postRouter.Use(authenticator.AuthenticateHttp, listingHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/listing/{id:%s}", regexUuid), listingHandler.HandlePutListing)
	putRouter.Use(authenticator.AuthenticateHttp, listingHandler.GetBody, listingHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/listing/{id:%s}", regexUuid), listingHandler.HandleDeleteListing)
	deleteRouter.Use(authenticator.AuthenticateHttp, listingHandler.GetIdFromPath)
}
