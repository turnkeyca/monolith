package user

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
	in := httptest.NewRequest(http.MethodPut, fmt.Sprintf("/api/user/%s", id.String()), nil)
	out := httptest.NewRecorder()
	ctx := context.WithValue(in.Context(), KeyBody{}, Dto{Id: id})
	ctx = context.WithValue(ctx, KeyId{}, id)
	in = in.WithContext(ctx)
	logger := log.New(os.Stdout, "", log.LstdFlags)
	db, _ := db.New(logger)
	handler := NewHandler(logger, db)
	handler.HandlePutUser(out, in)
	testQuery := db.GetNextTestQuery()
	assert.Equal(t, http.StatusNoContent, out.Code, "status code")
	assert.Equal(t, fmt.Sprintf("update user set id=%s, full_name=, email=, password=, phone_number=, nickname=, bio=, city=, province=, user_type=, send_notifications=, moving_reason=, has_roommates=, has_security_deposit=, is_smoker=, has_prev_lawsuit=, has_prev_eviction=, can_credit_check=, has_pets=, additional_details=0, move_in_date=, move_out_date=, property_management_company=, additional_details_lease=, monthly_budget_min=, monthly_budget_max= where id=%s;", id, id), testQuery, "body")
}
