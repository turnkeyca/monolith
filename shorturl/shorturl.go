package shorturl

import (
	"log"
	"net/http"

	"github.com/turnkeyca/monolith/bitly"
)

type Handler struct {
	logger      *log.Logger
	bitlyClient *bitly.Client
}

func NewHandler(logger *log.Logger, bitlyClient *bitly.Client) *Handler {
	return &Handler{
		logger:      logger,
		bitlyClient: bitlyClient,
	}
}

func (h *Handler) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/short-url", h.shortUrlHandler)
}

func (h *Handler) shortUrlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		h.getShortUrl(w, r)
		return
	}
	http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
}

func (h *Handler) getShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	shortUrl := New(h.bitlyClient.GetShortUrl((r.URL.Query().Get("url"))))
	err := shortUrl.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}
