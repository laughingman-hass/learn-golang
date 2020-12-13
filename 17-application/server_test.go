package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Hass":        20,
			"LaughingMan": 10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Hass' score", func(t *testing.T) {
		request := newGetScoreRequest("Hass")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertRespnseStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns LaughingMan's score", func(t *testing.T) {
		request := newGetScoreRequest("LaughingMan")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertRespnseStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Jinny")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertRespnseStatus(t, response.Code, http.StatusNotFound)
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

func assertRespnseStatus(t *testing.T, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("did not get the correct status, actual %d, expected %d", actual, expected)
	}
}

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
