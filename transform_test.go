package aural_test

import (
	"github.com/pspeter3/aural"
	"testing"
	"time"
)

func TestTransformArtists(t *testing.T) {
	followers := struct {
		Total uint64 `json:"total"`
	}{
		Total: 1,
	}
	externalURLs := struct {
		Spotify string `json:"spotify"`
	}{
		Spotify: "https://www.spotify.com",
	}
	empty := aural.Image{Height: 0, Width: 0, URL: ""}
	real := aural.Image{Height: 10, Width: 10, URL: "http://www.akamai.com"}
	artist := aural.Artist{
		Entity: aural.Entity{
			Name:         "test",
			ExternalURLs: externalURLs,
			Images: []aural.Image{
				empty,
				real,
			},
		},
		ID:         aural.ArtistID("1"),
		Genres:     []string{"foo", "bar"},
		Popularity: 1,
		Followers:  followers,
	}
	element := aural.TransformArtists([]aural.Artist{artist})[0]
	equals(t, element.Title, artist.Name)
	equals(t, element.ImageURL, real.URL)
	equals(t, element.Subtitle, "Popularity: 1 Followers: 1 Genres: foo,bar")
	buttons := []aural.Button{
		aural.NewWebButton("Open", externalURLs.Spotify),
		aural.NewPostbackButton("Top Tracks", "top-tracks:1"),
		aural.NewPostbackButton("Related Artists", "related-artists:1"),
	}
	for i, button := range buttons {
		equals(t, element.Buttons[i].Title, button.Title)
		equals(t, element.Buttons[i].Type, button.Type)
		equals(t, element.Buttons[i].URL, button.URL)
		equals(t, element.Buttons[i].Payload, button.Payload)
	}
}

func TestTransformTracks(t *testing.T) {
	externalURLs := struct {
		Spotify string `json:"spotify"`
	}{
		Spotify: "https://www.spotify.com",
	}
	empty := aural.Image{Height: 0, Width: 0, URL: ""}
	real := aural.Image{Height: 10, Width: 10, URL: "http://www.akamai.com"}
	track := aural.Track{
		Entity: aural.Entity{
			Name:         "test",
			ExternalURLs: externalURLs,
		},
		ID: aural.TrackID("1"),
		Album: aural.Entity{
			Name: "album",
			Images: []aural.Image{
				empty,
				real,
			},
		},
		Disc:     1,
		Track:    1,
		Duration: time.Second,
	}
	element := aural.TransformTracks([]aural.Track{track})[0]
	equals(t, element.Title, track.Name)
	equals(t, element.ImageURL, real.URL)
	equals(t, element.Subtitle, "album 1:1 1.00s")
	buttons := []aural.Button{
		aural.NewWebButton("Open", externalURLs.Spotify),
	}
	for i, button := range buttons {
		equals(t, element.Buttons[i].Title, button.Title)
		equals(t, element.Buttons[i].Type, button.Type)
		equals(t, element.Buttons[i].URL, button.URL)
		equals(t, element.Buttons[i].Payload, button.Payload)
	}
	track.Explicit = true
	element = aural.TransformTracks([]aural.Track{track})[0]
	equals(t, element.Subtitle, "album 1:1 1.00s Explicit")
}
