package core

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type BookmarkService interface {
	Create(context.Context, *Bookmark) error
	Get(context.Context, BookmarkId) (Bookmark, error)
	List(context.Context, BookmarkFilter, ...Opt) ([]Bookmark, error)
	Update(context.Context, BookmarkId, *Bookmark) error
	Delete(context.Context, BookmarkId) error

	GetBookmarkFromURL(context.Context, string) (Bookmark, error)
}

type BookmarkFilter interface {
	BookmarkFilter()
}

type Opt interface {
	Opt()
}

type BookmarkId uuid.UUID
type Bookmark struct {
	Id   uuid.UUID `json:"id,omitempty"`
	URL  string    `json:"url,omitempty"` //Must be unique
	Hash string    `json:"hash,omitempty"`
	Data struct {
		Title       string `json:"title,omitempty"`
		Type        string `json:"type,omitempty"`
		Url         string `json:"url,omitempty"`
		Site        string `json:"site,omitempty"`
		SiteName    string `json:"site_name,omitempty"`
		Description string `json:"description,omitempty"`
		Locale      string `json:"locale,omitempty"`
		Image       string `json:"image,omitempty"`
		Content     string `json:"content,omitempty"`
	} `json:"data,omitempty"`
}

func (l Bookmark) Valid() error {
	return nil
}
