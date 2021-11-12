package roommate

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/turnkeyca/monolith/auth"
	"github.com/turnkeyca/monolith/db"
	"github.com/turnkeyca/monolith/permission"
	"github.com/turnkeyca/monolith/util"
)

type Handler struct {
	logger     *log.Logger
	authorizer *permission.Authorizer
	db         *db.Database
}

func NewHandler(logger *log.Logger, db *db.Database, authorizer *permission.Authorizer) *Handler {
	return &Handler{
		logger:     logger,
		db:         db,
		authorizer: authorizer,
	}
}

type GenericError struct {
	Message string `json:"message"`
}

type ValidationError struct {
	Messages []string `json:"messages"`
}

func ConfigureRoommateRoutes(router *mux.Router, logger *log.Logger, database *db.Database, authenticator *auth.Authenticator, authorizer *permission.Authorizer) {
	roommateHandler := NewHandler(logger, database, authorizer)

	getRouter := router.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc(fmt.Sprintf("/v1/roommate/{id:%s}", util.REGEX_UUID), roommateHandler.HandleGetRoommate)
	getRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetIdFromPath, roommateHandler.CheckPermissionsRoommateIdView)
	getRouter.HandleFunc("/v1/roommate", roommateHandler.HandleGetRoommateByUserId)
	getRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetUserIdFromQueryParameters, roommateHandler.CheckPermissionsView)

	postRouter := router.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/v1/roommate", roommateHandler.HandlePostRoommate)
	postRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetBody, roommateHandler.CheckPermissionsBodyEdit)

	putRouter := router.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc(fmt.Sprintf("/v1/roommate/{id:%s}", util.REGEX_UUID), roommateHandler.HandlePutRoommate)
	putRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetBody, roommateHandler.GetIdFromPath, roommateHandler.CheckPermissionsRoommateIdEdit)

	deleteRouter := router.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc(fmt.Sprintf("/v1/roommate/{id:%s}", util.REGEX_UUID), roommateHandler.HandleDeleteRoommate)
	deleteRouter.Use(authenticator.AuthenticateHttp, roommateHandler.GetIdFromPath, roommateHandler.CheckPermissionsRoommateIdEdit)
}
