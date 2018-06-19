package webhook

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/nats-io/go-nats-streaming"
)

var (
	ErrWebhookNotActive = errors.New("client webhook is not active notifying sysadmin")
)

func Handler(client *http.Client, url string) func(*stan.Msg) {

	//stan.Msg is the message from the NATS durable queue
	fn := func(m *stan.Msg) {
		now := time.Now()
		var resp *http.Response

		//Unit of work
		fn := func() (ferr error) {
			resp, ferr = PostEvent(client, url, m.Data)
			return ferr
		}

		//Executes the unit of work and does a simple backoff
		if err := SimpleLinearBackOff(10, fn); err != nil {

			log.WithFields(log.Fields{
				"type": "consumer",
				"kind": "webhook",
			}).Error(err)

			return
		}

		log.WithFields(log.Fields{
			"type": "consumer",
			"kind": "webhook",
			"time": time.Since(now),
		}).Infof("Successful webhook [%s] - %s", resp.Status, resp.Request.URL)

	}

	return fn
}

func PostEvent(client *http.Client, url string, b []byte) (resp *http.Response, err error) {
	buf := bytes.NewBuffer(b)

	req, err := http.NewRequest(http.MethodPost, url, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Bookmarks API Server")
	req.Header.Set("Content-Type", "application/json")

	sigBuf := bytes.NewBuffer(buf.Bytes())
	//TODO Add valid Signature
	sigBuf.WriteString(":BookmarksSpecificSecretKeyHere")
	h := hmac.New(sha256.New, sigBuf.Bytes())
	req.Header.Set("Signature", base64.StdEncoding.EncodeToString(h.Sum(nil)))

	resp, err = client.Do(req)
	if err != nil {
		return
	}
	if resp.StatusCode >= 300 {
		err = ErrWebhookNotActive
		return
	}

	return
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
			"kind": "webhook",
		}).Errorf("WebhookHandler:: Retry count %d - %s", count, err)

		time.Sleep(time.Duration(count) * time.Second * 5)
	}

}
