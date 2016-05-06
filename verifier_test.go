package aural_test

import (
	"github.com/pspeter3/aural"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidToken(t *testing.T) {
	v := aural.NewVerifier("token")
	r, err := http.NewRequest(http.MethodGet, "/?hub.verify_token=evil", nil)
	ok(t, err)
	w := httptest.NewRecorder()
	v.ServeHTTP(w, r)
	equals(t, w.Code, http.StatusUnauthorized)
}

func TestValidToken(t *testing.T) {
	v := aural.NewVerifier("token")
	r, err := http.NewRequest(http.MethodGet, "/?hub.verify_token=token&hub.challenge=foo", nil)
	ok(t, err)
	w := httptest.NewRecorder()
	v.ServeHTTP(w, r)
	equals(t, w.Code, http.StatusOK)
	equals(t, w.Body.String(), "foo")
}
