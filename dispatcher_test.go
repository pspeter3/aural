package aural_test

import (
	"github.com/pspeter3/aural"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDispatcher_BadRequest(t *testing.T) {
	server := httptest.NewServer(aural.NewDispatcher(func() aural.Doer {
		return http.DefaultClient
	}, mockSender{}))
	defer server.Close()
	req, err := http.NewRequest(http.MethodPost, server.URL, strings.NewReader(""))
	ok(t, err)
	resp, err := http.DefaultClient.Do(req)
	ok(t, err)
	equals(t, resp.StatusCode, http.StatusBadRequest)
}

func TestDispatcher_ServeHTTP(t *testing.T) {
	user := aural.User{
		ID: aural.UserID(1),
	}
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Path, "/v1/search")
		equals(t, r.URL.Query().Get("q"), "Does Not Exist")
		equals(t, r.URL.Query().Get("type"), "artist")
		equals(t, r.URL.Query().Get("market"), "us")
		equals(t, r.URL.Query().Get("limit"), "10")
		w.Write([]byte(`{"artist":{"items":[]}}`))
	})}
	s := mockSender{t, user}
	server := httptest.NewServer(aural.NewDispatcher(func() aural.Doer {
		return m
	}, s))
	defer server.Close()
	req, err := http.NewRequest(http.MethodPost, server.URL, strings.NewReader(`{
		"entry": [{
			"messaging": [
				{
					"sender": {
						"id": 1
					},
					"message": {
						"text": "Does Not Exist"
					}
				}
			]
		}]
	}`))
	ok(t, err)
	resp, err := http.DefaultClient.Do(req)
	ok(t, err)
	equals(t, resp.StatusCode, http.StatusOK)
}
