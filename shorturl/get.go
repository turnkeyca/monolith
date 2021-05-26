package shorturl

import (
	"fmt"
	"net/http"
)

// swagger:route GET /api/short-url?longUrl={url} shorturl getShortUrl
// get a bitly short url

// Return a bitly short url
// responses:
//	200: shortUrlResponse
//	500: shortUrlErrorResponse

// HandleGetShortUrl handles GET requests
func (h *Handler) HandleGetShortUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := h.GetShortUrl(r.URL.Query().Get("url"))
	err := shortUrl.Write(w)
	if err != nil {
		http.Error(w, fmt.Sprintf("error creating short url: %#v\n", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) GetShortUrl(longUrl string) *ShortUrlDto {
	return New(h.bitlyClient.GetShortUrl(longUrl))
}
