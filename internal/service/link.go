package service

import (
	"context"

	"github.com/FourSigma/bookmarks/internal/core"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

func NewLinkService(db *sqlx.DB) core.LinkService {
	return linkService{db: db}
}

type linkService struct {
	db *sqlx.DB
}

func (l linkService) Create(ctx context.Context, link *core.Link) (err error) {
	if err = l.db.QueryRowContext(ctx, "INSERT INTO links(id, url, data) VALUES ($1, $2, $3) RETURNING *").Scan(link); err != nil {
		return
	}

	return
}
func (l linkService) Get(ctx context.Context, id core.LinkId) (link core.Link, err error) {
	if err = l.db.Get(&link, "SELECT * from links WHERE id = $1"); err != nil {
		return
	}
	return
}
func (l linkService) List(ctx context.Context, filt core.LinkFilter, opts ...core.Opt) (rs []core.Link, err error) {
	return l.list(ctx, filt, opts...)
}

func (l linkService) list(ctx context.Context, filt core.LinkFilter, opts ...core.Opt) (rs []core.Link, err error) {
	if err = l.db.Select(&rs, "SELECT * from links"); err != nil {
		return
	}
	return
}

func (l linkService) Update(ctx context.Context, id core.LinkId, modified *core.Link) (err error) {
	//	if err = l.db.ExecContext(ctx, "UPDATE ", args ...interface{})
	return
}

func (l linkService) Delete(ctx context.Context, id core.LinkId) (err error) {
	if _, err = l.db.ExecContext(ctx, "DELETE FROM links where id = $1", uuid.UUID(id)); err != nil {
		return
	}
	return
}

func (l linkService) GetLinkFromURL(ctx context.Context, url string) (link core.Link, err error) {
	return
}