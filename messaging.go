package aural

// Messaging is a struct received from Facebook Messenger
type Messaging struct {
	Sender  User `json:"sender"`
	Message struct {
		Text string `json:"text"`
	} `json:"message"`
	Postback struct {
		Payload string `json:"payload"`
	} `json:"postback"`
}
