package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/bitly"
	"github.com/turnkeyca/monolith/server"
	"github.com/turnkeyca/monolith/shorturl"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	err := godotenv.Load(".env")
	if err != nil {
		logger.Printf("failed to load environment from .env: %v", err)
	}
	mux := http.NewServeMux()

	shorturlHandler := shorturl.NewHandler(logger, bitly.NewClient(logger))
	shorturlHandler.SetupRoutes(mux)

	logger.Println("Starting server")
	srv := server.New(logger)
	err = srv.NewHttpServer(mux).ListenAndServeTLS("", "")
	if err != nil {
		logger.Fatalf("Failed to start %v", err)
	}
}
