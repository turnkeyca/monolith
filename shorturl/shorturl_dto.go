package shorturl

import (
	"encoding/json"
	"io"
)

type ShortUrlDto struct {
	Url string `json:"url"`
}

func New(url string) *ShortUrlDto {
	return &ShortUrlDto{
		Url: url,
	}
}

func (d *ShortUrlDto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}
