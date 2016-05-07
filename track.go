package aural

import (
	"time"
)

// TrackID is a type alias for track ids returned by the Spotify API
type TrackID string

// Track is a struct for tracks returned by the Spotify API
type Track struct {
	Entity
	ID         TrackID       `json:"id"`
	Album      Entity        `json:"album"`
	Duration   time.Duration `json:"duration_ms"`
	Explicit   bool          `json:"explicit"`
	Popularity uint64        `json:"popularity"`
}
