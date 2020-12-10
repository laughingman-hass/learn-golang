package context_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	data := "hello, world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	respnse := httptest.NewRecorder()

	svr.ServeHTTP(respnse, request)

	if respnse.Body.String() != data {
		t.Errorf(`actual "%s", expected "%s"`, respnse.Body.String(), data)
	}
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}
