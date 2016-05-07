package aural

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Sender is an interface for a Facebook Messenger sender
type Sender interface {
	Send(client Doer, user User, elements []Element) error
}

type sender struct {
	url string
}

type fbPayload struct {
	TemplateType string    `json:"template_type"`
	Elements     []Element `json:"elements"`
}

type fbAttachment struct {
	Type    string    `json:"type"`
	Payload fbPayload `json:"payload"`
}

type fbMessage struct {
	Attachment fbAttachment `json:"attachment"`
}

func (s sender) Send(client Doer, user User, elements []Element) error {
	body, err := json.Marshal(struct {
		Recipient User      `json:"recipient"`
		Message   fbMessage `json:"message"`
	}{
		Recipient: user,
		Message: fbMessage{
			Attachment: fbAttachment{
				Type: "template",
				Payload: fbPayload{
					TemplateType: "generic",
					Elements:     elements,
				},
			},
		},
	})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, s.url, bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	return err
}

// NewSender creates a new Facebook Messenger sender
func NewSender(token string) Sender {
	return &sender{"https://graph.facebook.com/v2.6/me/messages?access_token=" + token}
}
