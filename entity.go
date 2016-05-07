package aural

// Entity is the base type for Spotify records.
type Entity struct {
	Name         string `json:"name"`
	ExternalURLs struct {
		Spotify string `json:"spotify"`
	} `json:"external_urls"`
	Images []Image `json:"images"`
}
