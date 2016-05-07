package aural

// Button is a struct for a response in Facebook Messenger
type Button struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	URL     string `json:"url,omitempty"`
	Payload string `json:"payload,omitempty"`
}

// NewWebButton is a convenience constructor for a web_url button
func NewWebButton(title string, url string) Button {
	return Button{
		Type:  "web_url",
		Title: title,
		URL:   url,
	}
}

// NewPostbackButton is a convenience constructor for a postback button
func NewPostbackButton(title string, payload string) Button {
	return Button{
		Type:    "postback",
		Title:   title,
		Payload: payload,
	}
}
