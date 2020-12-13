package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Hass' score", func(t *testing.T) {
		request := newGetScoreRequest("Hass")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns LaughingMan's score", func(t *testing.T) {
		request := newGetScoreRequest("LaughingMan")
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetScoreRequest(name string) *http.Request {
	requst, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return requst
}

func assertResponseBody(t *testing.T, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("response body is incorrect, actual %q, expected %q", actual, expected)
	}
}
