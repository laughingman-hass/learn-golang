package facest

import "net/http"

type Client struct {
	BaseURL    string
	apiKey     string
	HTTPClient *http.Client
}
