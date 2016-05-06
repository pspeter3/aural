package aural

// Image is a struct returned by the Spotify API
type Image struct {
	Height uint64 `json:"height"`
	Width  uint64 `json:"width"`
	URL    string `json:"url"`
}
