package shorturl

import (
	"log"

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
