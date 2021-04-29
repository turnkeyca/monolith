package shorturl

import (
	"fmt"
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
	h.logger.Println("new short url request")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(h.getShortUrl(r.URL.Query().Get("url"))))
}

func (h *Handler) getShortUrl(url string) string {
	h.logger.Println(fmt.Sprintf("converting url %s", url))
	return h.bitlyClient.GetShortUrl(url)
}
