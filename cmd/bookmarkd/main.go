package main

import (
	"net/http"

	"github.com/FourSigma/bookmarks/internal/api"
	"github.com/FourSigma/bookmarks/internal/service"
	"github.com/FourSigma/bookmarks/internal/service/nats"
	"github.com/FourSigma/bookmarks/internal/service/nats/slack"
	"github.com/FourSigma/bookmarks/internal/service/nats/webhook"
	"github.com/FourSigma/bookmarks/pkg/log"
)

func main() {

	/***********************************
	   Webhook consumer Go routine
	************************************/
	wsc, err := nats.Subscribe(
		service.SubjectBookmarkCreated,
		webhook.Handler(http.DefaultClient, "https://webhook.site/dba8491d-84e3-4ca5-b25a-6a75cc1ecfe7"),
		"webhook",
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer wsc.Close()

	/***********************************
	   Slack consumer Go routine
	************************************/
	ssc, err := nats.Subscribe(
		service.SubjectBookmarkCreated,
		slack.Handler(
			"token",
			"bookmarks",
		),
		"slack",
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer ssc.Close()

	/***********************************
	  API Server Go routine
	************************************/
	log.Errorln(api.ListenAndServe("8080"))
}
