package user

import (
	"log"
	"net/http"

	"github.com/turnkeyca/monolith/db"
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

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/user", h.userHandler)
}

func (h *Handler) userHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		return
	}
	if r.Method == http.MethodPost {
		return
	}
	if r.Method == http.MethodPut {
		return
	}
	if r.Method == http.MethodDelete {
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
}
