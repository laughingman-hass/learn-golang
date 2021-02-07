package facest

import (
	"net/http"
	"time"
)

const (
	BaseURLV1 = "https://api.facest.io/v1"
)

func NewClient(appKey string) *Client {
	return &Client{
		BaseURL: BaseURLV1,
		apiKey:  appKey,
		HTTPClient: &http.Client{
			Timeout: time.Minute,
		},
	}
}

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}
