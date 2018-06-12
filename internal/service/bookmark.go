package service

import (
	"context"

	"github.com/FourSigma/bookmarks/internal/core"
	"github.com/FourSigma/bookmarks/pkg/opengraph"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
)

func NewBookmarkService(db *sqlx.DB, og opengraph.OGClient) core.BookmarkService {
	return linkService{
		db: db,
		og: og,
	}
}

type linkService struct {
	db *sqlx.DB
	og opengraph.OGClient
}

func (l linkService) Create(ctx context.Context, link *core.Bookmark) (err error) {
	if link.Id, err = uuid.NewV4(); err != nil {
		return
	}
	if err = l.db.QueryRowContext(ctx, "INSERT INTO links(id, url, data) VALUES ($1, $2, $3) RETURNING *").Scan(link); err != nil {
		return
	}

	return
}
func (l linkService) Get(ctx context.Context, id core.BookmarkId) (link core.Bookmark, err error) {
	if err = l.db.Get(&link, "SELECT * from links WHERE id = $1"); err != nil {
		return
	}
	return
}
func (l linkService) List(ctx context.Context, filt core.BookmarkFilter, opts ...core.Opt) (rs []core.Bookmark, err error) {
	return l.list(ctx, filt, opts...)
}

func (l linkService) list(ctx context.Context, filt core.BookmarkFilter, opts ...core.Opt) (rs []core.Bookmark, err error) {
	if err = l.db.Select(&rs, "SELECT * from links"); err != nil {
		return
	}
	return
}

func (l linkService) Update(ctx context.Context, id core.BookmarkId, modified *core.Bookmark) (err error) {
	//	if err = l.db.ExecContext(ctx, "UPDATE ", args ...interface{})
	return
}

func (l linkService) Delete(ctx context.Context, id core.BookmarkId) (err error) {
	if _, err = l.db.ExecContext(ctx, "DELETE FROM links where id = $1", uuid.UUID(id)); err != nil {
		return
	}
	return
}

func (l linkService) GetBookmarkFromURL(ctx context.Context, url string) (link core.Bookmark, err error) {
	m, err := l.og.OpenGraphMetaData(ctx, url)
	if err != nil {
		return
	}

	if val, ok := m["title"]; ok {
		link.Data.Title = val
	}

	if val, ok := m["type"]; ok {
		link.Data.Type = val
	}

	if val, ok := m["site_name"]; ok {
		link.Data.SiteName = val
	}

	if val, ok := m["description"]; ok {
		link.Data.Description = val
	}

	if val, ok := m["url"]; ok {
		link.Data.Url = val
	}

	if val, ok := m["image"]; ok {
		link.Data.Image = val
	}

	return
}
