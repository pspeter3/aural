package aural

import (
	"encoding/json"
	"log"
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
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("%v\n", data)
	var wg sync.WaitGroup
	for _, entry := range data.Entries {
		for _, messaging := range entry.Messagings {
			go func(client Doer, sender Sender, messaging Messaging) {
				wg.Add(1)
				err := Process(client, sender, messaging)
				if err != nil {
					log.Println(err)
				}
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
