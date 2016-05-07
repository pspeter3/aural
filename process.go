package aural

import (
	"errors"
)

// Process processes a messaging and sends a response
func Process(client Doer, sender Sender, messaging Messaging) error {
	var elements []Element
	if messaging.Postback.Payload != "" {
		method, artist, err := parseAction(action(messaging.Postback.Payload))
		if err != nil {
			return err
		}
		switch method {
		case actionTopTracks:
			tracks, err := TopTracks(client, artist)
			if err != nil {
				return err
			}
			elements = TransformTracks(tracks)
			break
		case actionRelatedArtists:
			artists, err := RelatedArtists(client, artist)
			if err != nil {
				return err
			}
			elements = TransformArtists(artists)
			break
		default:
			return errors.New("Invalid Action")
		}
	} else {
		artists, err := SearchArtists(client, messaging.Message.Text)
		if err != nil {
			return err
		}
		elements = TransformArtists(artists)
	}
	return sender.Send(client, messaging.Sender, elements)
}
