package aural

import (
	"fmt"
	"strings"
)

// ImageURL finds the largest image URL
func ImageURL(images []Image) string {
	var url string
	var width, height uint64
	for _, image := range images {
		if image.Width > width && image.Height > height {
			width = image.Width
			height = image.Height
			url = image.URL
		}
	}
	return url
}

func actionButton(title string, method string, artist Artist) Button {
	return NewPostbackButton(title, string(newAction(method, artist.ID)))
}

// TransformArtists transforms Spotify Artists into Messenger Elements
func TransformArtists(artists []Artist) []Element {
	elements := make([]Element, len(artists))
	for i, artist := range artists {
		genres := strings.Join(artist.Genres, ",")
		subtitle := fmt.Sprintf("Popularity: %d Followers: %d Genres: %s", artist.Popularity, artist.Followers.Total, genres)
		elements[i] = Element{
			Title:    artist.Name,
			ImageURL: ImageURL(artist.Images),
			Subtitle: subtitle,
			Buttons: []Button{
				NewWebButton("Open", artist.ExternalURLs.Spotify),
				actionButton("Top Tracks", actionTopTracks, artist),
				actionButton("Related Artists", actionRelatedArtists, artist),
			},
		}
	}
	return elements
}

// TransformTracks transforms Spotify Tracks into Messenger Elements
func TransformTracks(tracks []Track) []Element {
	elements := make([]Element, len(tracks))
	for i, track := range tracks {
		explicit := ""
		if track.Explicit {
			explicit = " Explicit"
		}
		subtitle := fmt.Sprintf("%s %d:%d %.2fs%s", track.Album.Name, track.Disc, track.Track, track.Duration.Seconds(), explicit)
		elements[i] = Element{
			Title:    track.Name,
			ImageURL: ImageURL(track.Album.Images),
			Subtitle: subtitle,
			Buttons: []Button{
				NewWebButton("Open", track.ExternalURLs.Spotify),
			},
		}
	}
	return elements
}
