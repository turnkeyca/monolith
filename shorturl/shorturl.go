package shorturl

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/bitly"
)

type Handler struct {
	logger      *log.Logger
	bitlyClient *bitly.Client
}

func NewHandler(logger *log.Logger, bitlyClient *bitly.Client) *Handler {
	return &Handler{
		logger:      logger,
		bitlyClient: bitlyClient,
	}
}

type GenericError struct {
	Message string `json:"message"`
}

func ConfigureShortUrlRoutes(router *mux.Router, logger *log.Logger, bitly *bitly.Client, authenticator *auth.Authenticator) {
	shorturlHandler := NewHandler(logger, bitly)
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/short-url", shorturlHandler.HandleGetShortUrl)
	getRouter.Use(authenticator.AuthenticateHttp)
}
