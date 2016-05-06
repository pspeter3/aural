package aural

// ArtistID is a type alias for artist ids returned by the Spotify API
type ArtistID string

// Artist is a struct for artists returned by the Spotify API
type Artist struct {
	Entity
	ID         ArtistID `json:"id"`
	Genres     []string `json:"genres"`
	Popularity uint64   `json:"popularity"`
	Followers  struct {
		Total uint64 `json:"total"`
	} `json:"followers"`
}
