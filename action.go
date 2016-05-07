package aural

import (
	"errors"
	"strings"
)

type action string

const actionSeparator = ":"
const actionTopTracks = "top-tracks"
const actionRelatedArtists = "related-artists"

func newAction(method string, artist ArtistID) action {
	return action(strings.Join([]string{method, string(artist)}, actionSeparator))
}

func parseAction(action action) (method string, artist ArtistID, err error) {
	parts := strings.Split(string(action), actionSeparator)
	if len(parts) != 2 {
		return "", "", errors.New("Invalid Action Format")
	}
	method = parts[0]
	artist = ArtistID(parts[1])
	return method, artist, nil
}
