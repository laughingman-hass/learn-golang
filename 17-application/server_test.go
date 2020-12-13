package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Hass' score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Hass", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		actual := response.Body.String()
		expected := "20"

		if actual != expected {
			t.Errorf("actual %q, exected %q", actual, expected)
		}
	})

	t.Run("returns LaughingMan's score", func(t *testing.T) {
        request, _ := http.NewRequest(http.MethodGet, "/players/LaughingMan", nil)
        response := httptest.NewRecorder()

        PlayerServer(response, request)

        actual := response.Body.String()
        expected := "10"

        if actual != expected {
            t.Errorf("actual %q, expected %q", actual, expected)
        }
	})
}
