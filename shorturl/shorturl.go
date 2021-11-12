package shorturl

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/authenticator"
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

func ConfigureShortUrlRoutes(router *mux.Router, logger *log.Logger, bitly *bitly.Client, authenticator *authenticator.Authenticator) {
	shorturlHandler := NewHandler(logger, bitly)
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/v1/shorturl", shorturlHandler.HandleGetShortUrl)
	getRouter.Use(authenticator.AuthenticateHttp)
}
