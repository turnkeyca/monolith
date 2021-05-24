package shorturl

import (
	"encoding/json"
	"io"
)

type Dto struct {
	Url string `json:"url"`
}

func New(url string) *Dto {
	return &Dto{
		Url: url,
	}
}

func (d *Dto) Write(w io.Writer) error {
	return json.NewEncoder(w).Encode(d)
}
