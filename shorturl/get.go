package shorturl

import "net/http"

func (h *Handler) HandleGetShortUrl(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	shortUrl := h.GetShortUrl(r.URL.Query().Get("url"))
	err := shortUrl.Write(w)
	if err != nil {
		h.logger.Printf("encoding error: %#v", err)
	}
}

func (h *Handler) GetShortUrl(longUrl string) *Dto {
	return New(h.bitlyClient.GetShortUrl(longUrl))
}
