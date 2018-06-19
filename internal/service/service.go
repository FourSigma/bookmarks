package service

import (
	"context"

	"github.com/FourSigma/bookmarks/pkg/log"
	"github.com/jmoiron/sqlx"
)

func GetDatabaseConnection() *sqlx.DB {
	db, err := sqlx.Connect("postgres", "user=web dbname=bookmarks password=letmein sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	log.Infoln("Connected to PostgreSQL database: bookmarks")
	return db
}

type Notifier interface {
	Notify(context.Context, interface{})
}

func NewDummyNotifier() Notifier {
	return dummyNotify{}
}

type dummyNotify struct{}

func (d dummyNotify) Notify(context.Context, interface{}) {
	return
}
