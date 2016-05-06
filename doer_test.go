package aural_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type mockDoer struct {
	handler http.HandlerFunc
}

func (m *mockDoer) Do(req *http.Request) (*http.Response, error) {
	server := httptest.NewServer(m.handler)
	transport := &http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}
	client := &http.Client{Transport: transport}
	return client.Do(req)
}

func TestMockDo(t *testing.T) {
	// Use a status code that does not exist in the wild
	code := 418
	m := mockDoer{
		handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(code)
		}),
	}
	req, err := http.NewRequest(http.MethodGet, "http://www.example.com", nil)
	ok(t, err)
	resp, err := m.Do(req)
	ok(t, err)
	equals(t, resp.StatusCode, code)
}
