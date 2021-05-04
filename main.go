package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/bitly"
	"github.com/turnkeyca/monolith/server"
	"github.com/turnkeyca/monolith/shorturl"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	err := godotenv.Load(".env")
	if err != nil {
		logger.Printf("failed to load environment from .env: %#v\n", err)
	}
	mux := http.NewServeMux()

	shorturlHandler := shorturl.NewHandler(logger, bitly.NewClient(logger))
	shorturlHandler.SetupRoutes(mux)

	logger.Println("Starting server")
	srv := server.New(logger)
	httpServer := srv.NewHttpServer(mux)
	go func() {
		err = httpServer.ListenAndServeTLS("", "")
		if err != nil {
			logger.Fatalf("Failed to start %#v", err)
		}
	}()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt)
	s := <-sc
	logger.Printf("termination signal received - trying to shutdown gracefully: %#v\n", s)
	c, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	cancelFunc()
	httpServer.Shutdown(c)
}
