package employment

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
	in := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/employment/%s", id.String()), nil)
	out := httptest.NewRecorder()
	ctx := context.WithValue(in.Context(), KeyBody{}, Dto{Id: id})
	ctx = context.WithValue(ctx, KeyId{}, id)
	in = in.WithContext(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, _, _ := db.New(logger)
	handler := NewHandler(logger, db)
	handler.HandlePutEmployment(out, in)
	testQuery := db.GetNextTestQuery()
	assert.Equal(t, http.StatusNoContent, out.Code, "status code")
	assert.Equal(t, fmt.Sprintf("update employments set id=%s where id=%s;", id, id), testQuery, "body")
}
