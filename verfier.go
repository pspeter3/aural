package aural

import (
	"net/http"
)

type verifier struct {
	token string
}

func (v verifier) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("hub.verify_token") == v.token {
		w.Write([]byte(r.URL.Query().Get("hub.challenge")))
	} else {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

// NewVerifier creates a Facebook Messenger token verifier
func NewVerifier(token string) http.Handler {
	return &verifier{token}
}
