package aural

import (
	"net/http"
)

type aural struct {
	dispatcher http.Handler
	verifier   http.Handler
}

func (a aural) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		a.verifier.ServeHTTP(w, r)
		break
	case http.MethodPost:
		a.dispatcher.ServeHTTP(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		break
	}
}

// New creates a new aural application
func New(dispatcher http.Handler, verifier http.Handler) http.Handler {
	return &aural{dispatcher, verifier}
}
