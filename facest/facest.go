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

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
