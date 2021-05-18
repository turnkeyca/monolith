package employment

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
	userId := uuid.MustParse("00b911af-6b87-4e68-9493-77a79bf8ccf2")
	in := httptest.NewRequest(http.MethodPost, "/api/employment", nil)
	out := httptest.NewRecorder()
	ctx := context.WithValue(in.Context(), KeyBody{}, Dto{UserId: &userId})
	in = in.WithContext(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, _ := db.New(logger)
	handler := NewHandler(logger, db)
	handler.HandlePostEmployment(out, in)
	assert.Equal(t, http.StatusNoContent, out.Code, "status code")
}
