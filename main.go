package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/bitly"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/server"
	"github.com/turnkeyca/monolith/shorturl"
	"github.com/turnkeyca/monolith/user"
)

const REGEX_UUID = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"

func configureDocRoutes(router *mux.Router) {
	getRouter := router.Methods(http.MethodGet).Subrouter()
	opts := middleware.RedocOpts{SpecURL: "./swagger.yml"}
	getRouter.Handle("/doc", middleware.Redoc(opts, nil))
	getRouter.Handle("/swagger.yml", http.FileServer(http.Dir("./")))
}

func configureShortUrlRoutes(router *mux.Router, logger *log.Logger, bitly *bitly.Client, authenticator *auth.Authenticator) {
	shorturlHandler := shorturl.NewHandler(logger, bitly)
	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/short-url", shorturlHandler.HandleGetShortUrl)
	getRouter.Use(authenticator.AuthenticateHttp)
}

func configureUserRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	userHandler := user.NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", REGEX_UUID), userHandler.HandleGetUser)
	getRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/user", userHandler.HandlePostUser)
	postRouter.Use(authenticator.AuthenticateHttp, userHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", REGEX_UUID), userHandler.HandlePutUser)
	putRouter.Use(authenticator.AuthenticateHttp, userHandler.GetBody, userHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/user/{id:%s}", REGEX_UUID), userHandler.HandleDeleteUser)
	deleteRouter.Use(authenticator.AuthenticateHttp, userHandler.GetIdFromPath)

}

func configureRoutes(logger *log.Logger) (*mux.Router, error) {
	router := mux.NewRouter()
	authenticator := auth.New(logger)
	bitly := bitly.NewClient(logger)
	database, errOpen, errPing := db.New(logger)
	if errOpen != nil {
		return nil, errOpen
	}
	if errPing != nil {
		return nil, errPing
	}

	configureDocRoutes(router)
	configureShortUrlRoutes(router, logger, bitly, authenticator)
	configureUserRoutes(router, logger, database, authenticator)

	return router, nil
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
	sm, err := configureRoutes(logger)
	if err != nil {
		logger.Fatalf("failed to configure routes: %#v\n", err)
	}
	logger.Println("starting server")
	srv := server.New(logger)
	httpServer := srv.NewHttpServer(sm)
	go serve(logger, httpServer)
	sc := make(chan os.Signal, 1)
	shutdown(logger, httpServer, sc)
}
