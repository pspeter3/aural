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

type mockRoundTripper struct {
	url *url.URL
}

func (m mockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = m.url.Scheme
	req.URL.Host = m.url.Host
	return http.DefaultTransport.RoundTrip(req)
}

func (m mockDoer) Do(req *http.Request) (*http.Response, error) {
	server := httptest.NewServer(m.handler)
	defer server.Close()
	serverURL, err := url.Parse(server.URL)
	if err != nil {
		return nil, err
	}
	client := &http.Client{Transport: mockRoundTripper{serverURL}}
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
	req, err := http.NewRequest(http.MethodGet, "https://www.example.com", nil)
	ok(t, err)
	resp, err := m.Do(req)
	ok(t, err)
	equals(t, resp.StatusCode, code)
}
