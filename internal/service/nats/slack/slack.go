package slack

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/FourSigma/bookmarks/internal/core"
	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/nlopes/slack"

	"github.com/nats-io/go-nats-streaming"
)

func Handler(token string, channel string) func(*stan.Msg) {
	now := time.Now()

	api := slack.New(token)

	fn := func(m *stan.Msg) {
		b := &core.Bookmark{}
		if err := json.Unmarshal(m.Data, b); err != nil {
			log.WithFields(log.Fields{
				"type": "consumer",
				"kind": "slack",
			}).Error(err)
		}

		params := slack.PostMessageParameters{}
		attachment := slack.Attachment{
			Title:     b.Data.Title,
			TitleLink: b.Data.Url,
			Pretext:   b.Data.SiteName,
			Text:      b.Data.Description,
			ThumbURL:  b.Data.Image,
		}
		params.Attachments = []slack.Attachment{attachment}

		//Unit of work
		fn := func() (ferr error) {
			_, _, ferr = api.PostMessage(channel,
				"",
				params,
			)
			return ferr
		}

		//Executes the unit of work and does a simple backoff
		if err := SimpleLinearBackOff(10, fn); err != nil {

			log.WithFields(log.Fields{
				"type": "consumer",
				"kind": "slack",
			}).Error(err)

			return
		}

		log.WithFields(log.Fields{
			"type": "consumer",
			"kind": "slack",
			"time": time.Since(now),
		}).Infoln("Successful slack post")

	}

	return fn
}

func SimpleLinearBackOff(max int64, fn func() error) error {
	var count int64 = 0
	for {

		err := fn()
		if err == nil {
			return nil
		}

		count = count + 1
		if count >= max {
			return errors.New("failed backoff")
		}

		log.WithFields(log.Fields{
			"type": "consumer",
			"kind": "slack",
		}).Errorf("SlackHandler:: Retry count %d - %s", count, err)

		time.Sleep(time.Duration(count) * time.Second * 5)
	}

}
