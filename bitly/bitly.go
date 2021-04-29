package bitly

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const BITLY_URL = "https://api-ssl.bitly.com/v4/shorten"
const BITLY_DOMAIN = "bit.ly"

type Client struct {
	logger *log.Logger
}

func NewClient(logger *log.Logger) *Client {
	return &Client{
		logger: logger,
	}
}

func (c *Client) GetShortUrl(longUrl string) string {
	requestBody, err := json.Marshal(map[string]string{
		"group_guid": os.Getenv("BITLY_GROUP_GUID"),
		"domain":     BITLY_DOMAIN,
		"long_url":   longUrl,
	})
	if err != nil {
		c.logger.Printf("json marshalling error occurred: %v", err)
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("POST", BITLY_URL, bytes.NewBuffer(requestBody))
	if err != nil {
		c.logger.Printf("creating request error occurred: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("BITLY_API_KEY")))
	resp, err := client.Do(req)
	if err != nil {
		c.logger.Printf("bit.ly error occurred: %v", err)
	}
	defer resp.Body.Close()
	var responseBody map[string]string
	json.NewDecoder(resp.Body).Decode(&responseBody)
	return responseBody["link"]
}
