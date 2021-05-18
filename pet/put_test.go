package pet

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

func TestPut(t *testing.T) {
	os.Setenv("TEST", "true")
	id := uuid.MustParse("ddeff7cc-da4c-4a26-b1fc-b023553abe82")
	userId := uuid.MustParse("00b911af-6b87-4e68-9493-77a79bf8ccf2")
	in := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/pet/%s", id.String()), nil)
	out := httptest.NewRecorder()
	ctx := context.WithValue(in.Context(), KeyBody{}, Dto{Id: id, UserId: &userId})
	ctx = context.WithValue(ctx, KeyId{}, id)
	in = in.WithContext(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, _ := db.New(logger)
	handler := NewHandler(logger, db)
	handler.HandlePutPet(out, in)
	testQuery := db.GetNextTestQuery()
	assert.Equal(t, http.StatusNoContent, out.Code, "status code")
	assert.Equal(t, fmt.Sprintf("update pet set id=%s, user_id=%s, breed=, weight=0.00 where id=%s;", id, userId, id), testQuery, "body")
}
