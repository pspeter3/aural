package aural

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
)

const spotifyLimit = "10"
const spotifyMarket = "us"

type artists struct {
	Artists []Artist `json:"artists"`
}

type tracks struct {
	Tracks []Track `json:"tracks"`
}

func spotify(client Doer, path string, values url.Values, data interface{}) error {
	spotifyURL := url.URL{
		Scheme:   "https",
		Host:     "api.spotify.com",
		Path:     "/v1/" + path,
		RawQuery: values.Encode(),
	}
	req, err := http.NewRequest(http.MethodGet, spotifyURL.String(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(data)
}

// SearchArtists queries Spotify for artists that match the query
func SearchArtists(client Doer, query string) ([]Artist, error) {
	values := url.Values{}
	values.Set("q", query)
	values.Set("type", "artist")
	values.Set("market", spotifyMarket)
	values.Set("limit", spotifyLimit)
	var data artists
	err := spotify(client, "search", values, &data)
	return data.Artists, err
}

// RelatedArtists queries Spotify for artists that are similar to artist
func RelatedArtists(client Doer, artist ArtistID) ([]Artist, error) {
	var data artists
	err := spotify(client, path.Join("artists", string(artist), "related-artists"), url.Values{}, &data)
	return data.Artists, err
}

// TopTracks queries Spotify for the top tracks of an artist
func TopTracks(client Doer, artist ArtistID) ([]Track, error) {
	values := url.Values{}
	values.Set("country", spotifyMarket)
	var data tracks
	err := spotify(client, path.Join("artists", string(artist), "top-tracks"), values, &data)
	return data.Tracks, err
}
