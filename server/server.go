package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

const CERT_FOLDER = ".certs"

type Server struct {
	logger *log.Logger
}

func New(logger *log.Logger) *Server {
	return &Server{
		logger: logger,
	}
}

func (s *Server) NewHttpServer(sm *mux.Router) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      sm,
		ErrorLog:     s.logger,
	}
}
