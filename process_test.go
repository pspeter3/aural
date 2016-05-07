package aural_test

import (
	"fmt"
	"github.com/pspeter3/aural"
	"net/http"
	"path"
	"strings"
	"testing"
)

type mockSender struct {
	t    *testing.T
	user aural.User
}

func (m mockSender) Send(client aural.Doer, user aural.User, elements []aural.Element) error {
	assert(m.t, strings.EqualFold(string(m.user.ID), string(user.ID)), "Ids must be equal")
	equals(m.t, len(elements), 0)
	return nil
}

func testAction(t *testing.T, method string, entityType string) {
	artist := "1"
	postback := struct {
		Payload string `json:"payload"`
	}{
		Payload: strings.Join([]string{method, artist}, ":"),
	}
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Path, "/"+path.Join("v1", "artists", artist, method))
		w.Write([]byte(fmt.Sprintf(`{"%s":{"items":[]}}`, entityType)))
	})}
	user := aural.User{
		ID: aural.UserID("1"),
	}
	s := mockSender{
		t:    t,
		user: user,
	}
	err := aural.Process(m, s, aural.Messaging{
		Sender:   user,
		Postback: postback,
	})
	ok(t, err)
}

func TestProcess_TopTracks(t *testing.T) {
	testAction(t, "top-tracks", "tracks")
}

func TestProcess_RelatedArtists(t *testing.T) {
	testAction(t, "related-artists", "artists")
}

func TestProcess_InvalidAction(t *testing.T) {
	postback := struct {
		Payload string `json:"payload"`
	}{
		Payload: "invalid",
	}
	err := aural.Process(nil, nil, aural.Messaging{
		Postback: postback,
	})
	assert(t, err != nil, "Must catch invalid action format")
	postback.Payload += ":1"
	err = aural.Process(nil, nil, aural.Messaging{
		Postback: postback,
	})
	assert(t, err != nil, "Must catch invalid action")
}

func TestProcess_SearchArtists(t *testing.T) {
	message := struct {
		Text string `json:"text"`
	}{
		Text: "Does Not Exist",
	}
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Query().Get("q"), message.Text)
		w.Write([]byte(`{"artist":{"items":[]}}`))
	})}
	user := aural.User{
		ID: aural.UserID("1"),
	}
	s := mockSender{
		t:    t,
		user: user,
	}
	err := aural.Process(m, s, aural.Messaging{
		Sender:  user,
		Message: message,
	})
	ok(t, err)
}
