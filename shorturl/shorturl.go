package shorturl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/turnkeyca/monolith/bitly"
)

type ShortUrlDto struct {
	Url string `json:"url"`
}

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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	shortUrl := ShortUrlDto{
		Url: h.getShortUrl(r.URL.Query().Get("url")),
	}
	resp, err := json.Marshal(shortUrl)
	if err != nil {
		h.logger.Printf("json marchalling error: %v", err)
		return
	}
	w.Write(resp)
}

func (h *Handler) getShortUrl(url string) string {
	h.logger.Println(fmt.Sprintf("converting url %s", url))
	return h.bitlyClient.GetShortUrl(url)
}
