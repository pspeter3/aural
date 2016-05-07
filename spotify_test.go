package aural_test

import (
	"github.com/pspeter3/aural"
	"net/http"
	"path"
	"testing"
)

func TestSearchArtists(t *testing.T) {
	query := "Does Not Exist"
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Path, "/v1/search")
		equals(t, r.URL.Query().Get("q"), query)
		equals(t, r.URL.Query().Get("type"), "artist")
		equals(t, r.URL.Query().Get("market"), "us")
		equals(t, r.URL.Query().Get("limit"), "10")
		w.Write([]byte(`{"artist":[]}`))
	})}
	artists, err := aural.SearchArtists(m, query)
	ok(t, err)
	equals(t, len(artists), 0)
}

func TestRelatedArtists(t *testing.T) {
	artist := aural.ArtistID("abcdef")
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Path, "/"+path.Join("v1", "artists", string(artist), "related-artists"))
		w.Write([]byte(`{"artist":[]}`))
	})}
	artists, err := aural.RelatedArtists(m, artist)
	ok(t, err)
	equals(t, len(artists), 0)
}

func TestTopTracks(t *testing.T) {
	artist := aural.ArtistID("abcdef")
	m := mockDoer{http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		equals(t, r.URL.Path, "/"+path.Join("v1", "artists", string(artist), "top-tracks"))
		equals(t, r.URL.Query().Get("country"), "us")
		w.Write([]byte(`{"tracks":[]}`))
	})}
	tracks, err := aural.TopTracks(m, artist)
	ok(t, err)
	equals(t, len(tracks), 0)
}
