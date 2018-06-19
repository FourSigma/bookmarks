package main

import (
	"net/http"

	"github.com/FourSigma/bookmarks/internal/api"
	"github.com/FourSigma/bookmarks/internal/service"
	"github.com/FourSigma/bookmarks/internal/service/nats/webhook"
	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/nats-io/go-nats-streaming"
)

func main() {

	/***********************************
	   Init Webhook consumer Go routine
	************************************/
	consumerWebhook, err := stan.Connect("foursigma-cluster", "consumer-webhook")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := consumerWebhook.Subscribe(
		service.SubjectBookmarkCreated, // bookmarks.events.create
		webhook.Handler(http.DefaultClient,
			"https://webhook.site/dba8491d-84e3-4ca5-b25a-6a75cc1ecfe7",
		), //Function closure func(*stan.Msg)
		stan.DurableName("webhook"), //Keeps track of the last consumed message offset by consumer
	); err != nil {
		log.Fatalln(err)
	}

	/***********************************
	  API Server Go routine
	************************************/
	log.Fatalln(api.ListenAndServe("8080"))
}
