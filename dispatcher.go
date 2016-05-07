package aural

import (
	"encoding/json"
	"net/http"
	"sync"
)

type dispatcher struct {
	newDoer func() Doer
	sender  Sender
}

func (d dispatcher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Entries []struct {
			Messagings []Messaging `json:"messaging"`
		} `json:"entry"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var wg sync.WaitGroup
	for _, entry := range data.Entries {
		for _, messaging := range entry.Messagings {
			go func(client Doer, sender Sender, messaging Messaging) {
				wg.Add(1)
				Process(client, sender, messaging)
				wg.Done()
			}(d.newDoer(), d.sender, messaging)
		}
	}
	wg.Wait()
}

// NewDispatcher creates a Messenger dispatcher
func NewDispatcher(newDoer func() Doer, sender Sender) http.Handler {
	return &dispatcher{newDoer, sender}
}
