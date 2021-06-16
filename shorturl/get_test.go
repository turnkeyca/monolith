package shorturl

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/turnkeyca/monolith/bitly"
)

func TestGetShortUrl(t *testing.T) {
	os.Setenv("TEST", "true")
	in := httptest.NewRequest("GET", "/api/short-url?url=blah", nil)
	out := httptest.NewRecorder()
	logger := log.New(os.Stdout, "", log.LstdFlags)
	shorturlHandler := NewHandler(logger, bitly.NewClient(logger))
	shorturlHandler.HandleGetShortUrl(out, in)
	if out.Code != http.StatusOK {
		t.Fatalf("expected: %d\tgot: %d", http.StatusOK, out.Code)
	}
	if out.Body.String() != "{\"url\":\"blah\"}\n" {
		t.Fatalf("expected: %s\tgot: %s", "{\"url\":\"blah\"}", out.Body.String())
	}
}
