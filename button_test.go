package aural_test

import (
	"github.com/pspeter3/aural"
	"testing"
)

func TestNewWebButton(t *testing.T) {
	title := "title"
	url := "https://www.example.com"
	button := aural.NewWebButton(title, url)
	equals(t, button.Type, "web_url")
	equals(t, button.Title, title)
	equals(t, button.URL, url)
}

func TestNewPostbackButton(t *testing.T) {
	title := "title"
	payload := "related-artists:1"
	button := aural.NewPostbackButton(title, payload)
	equals(t, button.Type, "postback")
	equals(t, button.Title, title)
	equals(t, button.Payload, payload)
}
