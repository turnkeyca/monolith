package shorturl

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetShortUrl(t *testing.T) {
	in := httptest.NewRequest("GET", "/api/short-url?url=blah", nil)
	out := httptest.NewRecorder()
	shorturlHandler := NewHandler(log.New(os.Stdout, "", log.LstdFlags))
	shorturlHandler.shortUrlHandler(out, in)
	if out.Code != http.StatusOK {
		t.Logf("expected: %d\tgot: %d", http.StatusOK, out.Code)
	}
	if out.Body.String() != "blah" {
		t.Logf("expected: %d\tgot: %d", http.StatusOK, out.Code)
	}
}
