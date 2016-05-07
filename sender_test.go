package aural_test

import (
	"github.com/pspeter3/aural"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestSender_Send(t *testing.T) {
	token := "token"
	sender := aural.NewSender(token)
	user := aural.User{
		ID: aural.UserID(1),
	}
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Path, "/v2.6/me/messages")
		equals(t, r.URL.Query().Get("access_token"), token)
		body, err := ioutil.ReadAll(r.Body)
		ok(t, err)
		expected := `{"recipient":{"id":1},"message":{"attachment":{"type":"template","payload":{"template_type":"generic","elements":[]}}}}`
		assert(t, strings.EqualFold(string(body), expected), "Bodies are not the same")
	})}
	sender.Send(m, user, []aural.Element{})
}
