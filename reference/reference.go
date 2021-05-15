package reference

import (
	"log"

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
