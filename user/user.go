package user

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
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

func (h *Handler) HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	user, err := h.GetUser(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting user by id: %s, %#v\n", id, err), http.StatusBadRequest)
	}
	err = user.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetUser(id uuid.UUID) (*Dto, error) {
	result, err := h.db.Query("select * from users where id = %s;", id.String())
	return Assemble(result), err
}

func (h *Handler) HandlePostUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	dto := r.Context().Value(KeyBody{}).(Dto)
	err := h.CreateUser(&dto)
	if err != nil {
		h.logger.Printf("saving error: %#v", err)
	}
	http.NoBody.WriteTo(w)
}

func (h *Handler) CreateUser(dto *Dto) error {
	err := h.db.Put("insert into users (id) values (%s);", dto.Id.String())
	return err
}

func (h *Handler) HandlePutUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	id := r.Context().Value(KeyId{}).(uuid.UUID)
	dto := r.Context().Value(KeyBody{}).(Dto)
	dto.Id = id
	err := h.UpdateUser(&dto)
	if err != nil {
		h.logger.Printf("saving error: %#v", err)
	}
	http.NoBody.WriteTo(w)
}

func (h *Handler) UpdateUser(dto *Dto) error {
	err := h.db.Put("update users set id=%s where id=%s", dto.Id.String(), dto.Id.String())
	return err
}

func (h *Handler) GetIdFromPath(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.MustParse(mux.Vars(r)["id"])
		ctx := context.WithValue(r.Context(), KeyId{}, id)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (h *Handler) GetBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		d, err := Read(r.Body)
		if err != nil {
			h.logger.Printf("decoding error: %#v", err)
			http.Error(w, "Error reading product", http.StatusBadRequest)
		}
		err = d.Validate()
		if err != nil {
			h.logger.Printf("validation error: %#v", err)
			http.Error(
				w,
				fmt.Sprintf("Error validating product: %s", err),
				http.StatusBadRequest,
			)
			return
		}
		ctx := context.WithValue(r.Context(), KeyBody{}, d)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
