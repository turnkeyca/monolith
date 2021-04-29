package main

import (
	"log"
	"net/http"
	"os"

	"github.com/turnkeyca/monolith/bitly"
	"github.com/turnkeyca/monolith/server"
	"github.com/turnkeyca/monolith/shorturl"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	mux := http.NewServeMux()

	shorturlHandler := shorturl.NewHandler(logger, bitly.NewClient(logger))
	shorturlHandler.SetupRoutes(mux)

	logger.Println("Starting server")
	err := server.New(mux).ListenAndServeTLS("", "")
	if err != nil {
		logger.Fatalf("Failed to start %v", err)
	}
}
