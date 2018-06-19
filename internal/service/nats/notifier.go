package nats

import (
	"context"
	"encoding/json"

	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/nats-io/go-nats-streaming"
)

func NewNATSConnection(clientId string) stan.Conn {
	conn, err := stan.Connect(
		"foursigma-cluster",
		clientId,
		stan.NatsURL(stan.DefaultNatsURL),
	)
	if err != nil {
		log.Fatalln("Unable to connect to NATS Server...", err)
	}

	log.WithFields(log.Fields{
		"type":     "NATS Connection",
		"clientId": clientId,
	}).Info("NATS connection established....")
	return conn
}

func NewNATSNotifier(topic string, conn stan.Conn) natsNotifier {
	return natsNotifier{
		topic: topic,
		sc:    conn,
	}
}

type natsNotifier struct {
	sc    stan.Conn
	topic string
}

func (n natsNotifier) Notify(ctx context.Context, v interface{}) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}

	if err = n.sc.Publish(n.topic, b); err != nil {
		log.Errorf("NATS publish error - %s", err)
		return
	}

}
