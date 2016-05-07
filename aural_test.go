package aural_test

import (
	"github.com/pspeter3/aural"
	"net/http"
	"net/http/httptest"
	"testing"
)

func assert(t *testing.T, condition bool, message string) {
	if !condition {
		t.Fatal(message)
	}
}

func ok(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

func equals(t *testing.T, actual, expected interface{}) {
	if actual != expected {
		t.Fatalf("%s does not equal %s", actual, expected)
	}
}

func TestAural(t *testing.T) {
	dispatcher := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	verifier := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusUnauthorized)
	})
	server := httptest.NewServer(aural.New(dispatcher, verifier))
	defer server.Close()
	tests := make(map[string]int)
	tests[http.MethodGet] = http.StatusUnauthorized
	tests[http.MethodPost] = http.StatusOK
	tests[http.MethodHead] = http.StatusMethodNotAllowed
	for method, status := range tests {
		req, err := http.NewRequest(method, server.URL, nil)
		ok(t, err)
		resp, err := http.DefaultClient.Do(req)
		ok(t, err)
		equals(t, resp.StatusCode, status)
	}
}
