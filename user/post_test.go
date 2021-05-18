package user

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/turnkeyca/monolith/db"
)

func TestPost(t *testing.T) {
	os.Setenv("TEST", "true")
	id := uuid.MustParse("ddeff7cc-da4c-4a26-b1fc-b023553abe82")
	in := httptest.NewRequest(http.MethodPost, "/api/user", nil)
	out := httptest.NewRecorder()
	ctx := context.WithValue(in.Context(), KeyBody{}, Dto{Id: id})
	in = in.WithContext(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, _ := db.New(logger)
	handler := NewHandler(logger, db)
	handler.HandlePostUser(out, in)
	assert.Equal(t, http.StatusNoContent, out.Code, "status code")
}
