package aural

import (
	"net/http"
)

// Doer is an interface for an object that can make HTTP requests
type Doer interface {
	Do(req *http.Request) (resp *http.Response, err error)
}
