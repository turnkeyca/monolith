package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/bitly"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/server"
	"github.com/turnkeyca/monolith/shorturl"
	"github.com/turnkeyca/monolith/user"
)

const REGEX_UUID = "[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}"

func configureShortUrlRoutes(router *mux.Router, logger *log.Logger, bitly *bitly.Client, authenticator *auth.Authenticator) {
	shorturlHandler := shorturl.NewHandler(logger, bitly)
	shortUrlGetRouter := router.Methods(http.MethodGet).Subrouter()
	shortUrlGetRouter.HandleFunc("/api/short-url", shorturlHandler.HandleGetShortUrl)
	shortUrlGetRouter.Use(authenticator.AuthenticateHttp)
}

func configureUserRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	userHandler := user.NewHandler(logger, database)

	userGetRouter := router.Methods(http.MethodGet).Subrouter()
	userGetRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", REGEX_UUID), userHandler.HandleGetUser)
	userGetRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath)

	userPostRouter := router.Methods(http.MethodPost).Subrouter()
	userPostRouter.HandleFunc("/api/user", userHandler.HandlePostUser)
	userPostRouter.Use(authenticator.AuthenticateHttp, userHandler.GetBody)

	userPutRouter := router.Methods(http.MethodPost).Subrouter()
	userPutRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", REGEX_UUID), userHandler.HandlePutUser)
	userPutRouter.Use(authenticator.AuthenticateHttp, userHandler.GetBody, userHandler.GetIdFromPath)
}

func configureRoutes(logger *log.Logger) *mux.Router {
	router := mux.NewRouter()
	authenticator := auth.New(logger)
	database := db.New(logger)
	bitly := bitly.NewClient(logger)

	configureShortUrlRoutes(router, logger, bitly, authenticator)
	configureUserRoutes(router, logger, database, authenticator)

	return router
}

func serve(logger *log.Logger, s *http.Server) {
	err := s.ListenAndServeTLS("", "")
	if err != nil {
		logger.Fatalf("failed to start %#v\n", err)
	}
}

func shutdown(logger *log.Logger, httpServer *http.Server, sc chan os.Signal) {
	signal.Notify(sc, os.Interrupt)
	s := <-sc
	logger.Printf("termination signal received - trying to shutdown gracefully: %#v\n", s)
	c, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	cancelFunc()
	httpServer.Shutdown(c)
}

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	err := godotenv.Load(".env")
	if err != nil {
		logger.Printf("failed to load environment from .env: %#v\n", err)
	}
	sm := configureRoutes(logger)

	logger.Println("starting server")
	srv := server.New(logger)
	httpServer := srv.NewHttpServer(sm)
	go serve(logger, httpServer)
	sc := make(chan os.Signal, 1)
	shutdown(logger, httpServer, sc)
}
