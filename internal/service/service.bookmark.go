package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"database/sql"

	"github.com/FourSigma/bookmarks/internal/core"
	"github.com/FourSigma/bookmarks/internal/service/nats"
	"github.com/FourSigma/bookmarks/pkg/opengraph"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

const SubjectBookmarkCreated = "bookmarks.event.create"

func NewBookmarkService(db *sqlx.DB, og opengraph.OGClient) core.BookmarkService {
	return bookmarkService{
		db:     db,
		og:     og,
		notify: nats.NewNATSNotifier(SubjectBookmarkCreated, nats.NewNATSConnection("publisher")),
	}
}

type bookmarkService struct {
	db     *sqlx.DB
	og     opengraph.OGClient
	notify Notifier
}

func (l bookmarkService) Create(ctx context.Context, bookmark *core.Bookmark) (err error) {
	if bookmark.Id, err = uuid.NewV4(); err != nil {
		return
	}
	bookmark.URL = bookmark.Data.Url
	if _, err = l.db.ExecContext(ctx, "INSERT INTO bookmarks(id, url, data) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", bookmark.Id, bookmark.URL, bookmark.Data); err != nil {
		return
	}

	l.notify.Notify(ctx, bookmark)
	return
}

func (l bookmarkService) Get(ctx context.Context, id core.BookmarkId) (bookmark core.Bookmark, err error) {
	if err = l.db.Get(&bookmark, "SELECT * from bookmarks WHERE id = $1"); err != nil {
		return
	}
	return
}

func (l bookmarkService) List(ctx context.Context, filt core.BookmarkFilter, opts ...core.Opt) (rs []core.Bookmark, err error) {
	return l.list(ctx, filt, opts...)
}

func (l bookmarkService) list(ctx context.Context, filt core.BookmarkFilter, opts ...core.Opt) (rs []core.Bookmark, err error) {
	rs = []core.Bookmark{}

	var query string
	var args []interface{}
	switch filt.(type) {
	case core.FilterBookmarksAll:
		query = "SELECT * from bookmarks"
	default:
		return nil, errors.New("unknown bookmarks filter")
	}

	if err = l.db.Select(&rs, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return []core.Bookmark{}, nil
		}
		fmt.Println(err)
		return
	}
	return
}

func (l bookmarkService) Update(ctx context.Context, id core.BookmarkId, modified *core.Bookmark) (err error) {
	data, err := json.Marshal(modified.Data)
	if err != nil {
		return
	}
	if _, err = l.db.ExecContext(ctx, "UPDATE bookmarks SET url=$1, SET data=$2", modified.URL, data); err != nil {
		return
	}
	return
}

func (l bookmarkService) Delete(ctx context.Context, id core.BookmarkId) (err error) {
	if _, err = l.db.ExecContext(ctx, "DELETE FROM bookmarks where id = $1", uuid.UUID(id)); err != nil {
		return
	}
	return
}

func (l bookmarkService) GetBookmarkFromURL(ctx context.Context, url string) (bookmark core.Bookmark, err error) {
	m, err := l.og.OpenGraphMetaData(ctx, url)
	if err != nil {
		return
	}

	if val, ok := m["title"]; ok {
		bookmark.Data.Title = val
	}

	if val, ok := m["type"]; ok {
		bookmark.Data.Type = val
	}

	if val, ok := m["site_name"]; ok {
		bookmark.Data.SiteName = val
	}

	if val, ok := m["description"]; ok {
		bookmark.Data.Description = val
	}

	if val, ok := m["url"]; ok {
		bookmark.Data.Url = val
	}

	if val, ok := m["image"]; ok {
		bookmark.Data.Image = val
	}

	return
}
