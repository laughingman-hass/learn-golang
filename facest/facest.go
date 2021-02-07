package facest

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
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

func (c *Client) sendRequest(ctx context.Context, req *http.Request, v interface{}) error {
	req = req.WithContext(ctx)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		var errRes errorResponse
		if err := json.NewDecoder(res.Body).Decode(&errRes); err != nil {
			return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
		}

		return errors.New(errRes.Message)
	}

	fullResponse := successResponse{
		Data: v,
	}
	if err := json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
		return err
	}

	return nil
}
