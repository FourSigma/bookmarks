package nats

import (
	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/nats-io/go-nats-streaming"
)

func Subscribe(subject string, h func(*stan.Msg), name string) (sc stan.Conn, err error) {
	sc, err = stan.Connect("foursigma-cluster", "consumer-"+name)
	if err != nil {
		log.Errorln(err)
		return
	}

	if _, err = sc.Subscribe(subject, h, stan.DurableName(name)); err != nil {
		log.Errorln(err)
		return
	}

	return
}
