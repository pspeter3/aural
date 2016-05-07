package main

import (
	"flag"
	"github.com/pspeter3/aural"
	"log"
	"net/http"
)

func main() {
	access := flag.String("access", "", "Facebook Access Token")
	addr := flag.String("addr", ":3000", "Address to bind")
	verify := flag.String("verify", "", "Facebook Verify Token")
	flag.Parse()
	sender := aural.NewSender(access)
	dispatcher := aural.NewDispatcher(func() aural.Doer {
		return &http.Client{}
	}, sender)
	verifier := aural.NewVerifier(verify)
	log.Fatal(http.ListenAndServe(*addr, aural.New(dispatcher, verifier)))
}
