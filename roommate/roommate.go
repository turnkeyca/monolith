package roommate

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/util"
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

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func ConfigureRoommateRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator) {
	roommateHandler := NewHandler(logger, database)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/api/roommate/{id:%s}", util.REGEX_UUID), roommateHandler.HandleGetRoommate)
	getRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetIdFromPath)
	getRouter.HandleFunc("/api/roommate", roommateHandler.HandleGetRoommateByUserId)
	getRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetUserIdFromQueryParameters)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/roommate", roommateHandler.HandlePostRoommate)
	postRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetBody)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/api/roommate/{id:%s}", util.REGEX_UUID), roommateHandler.HandlePutRoommate)
	putRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetBody, roommateHandler.GetIdFromPath)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/api/roommate/{id:%s}", util.REGEX_UUID), roommateHandler.HandleDeleteRoommate)
	deleteRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetIdFromPath)
}
