package listing

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/turnkeyca/monolith/db"
)

func TestGet(t *testing.T) {
	os.Setenv("TEST", "true")
	id := uuid.MustParse("ddeff7cc-da4c-4a26-b1fc-b023553abe82")
	in := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/listing/%s", id.String()), nil)
	out := httptest.NewRecorder()
	ctx := context.WithValue(in.Context(), KeyId{}, id)
	in = in.WithContext(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, _ := db.New(logger)
	dto := []interface{}{Dto{Id: id}}
	db.SetNextTestReturn(dto)
	handler := NewHandler(logger, db)
	handler.HandleGetListing(out, in)
	assert.Equal(t, http.StatusOK, out.Code, "status code")
	assert.Equal(t, fmt.Sprintf("{\"id\":\"%s\",\"fullName\":\"billy\"}\n", id), out.Body.String(), "body")
}
