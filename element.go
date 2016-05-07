package aural

// Element is a struct for an element that can be returned to Facebook Messenger
type Element struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	Subtitle string   `json:"subtitle"`
	Buttons  []Button `json:"buttons"`
}
