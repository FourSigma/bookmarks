package core

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

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
	Id  uuid.UUID `json:"id,omitempty"`
	URL string    `json:"url,omitempty"` //Must be unique
	//Hash string    `json:"hash,omitempty"`
	Data OpenGraphMetaData `json:"data,omitempty"`
}

func (l Bookmark) Valid() error {
	return nil
}

type OpenGraphMetaData struct {
	Title       string `json:"title,omitempty"`
	Type        string `json:"type,omitempty"`
	Url         string `json:"url,omitempty"`
	Site        string `json:"site,omitempty"`
	SiteName    string `json:"site_name,omitempty"`
	Description string `json:"description,omitempty"`
	Locale      string `json:"locale,omitempty"`
	Image       string `json:"image,omitempty"`
	Content     string `json:"content,omitempty"`
}

func (o OpenGraphMetaData) Value() (driver.Value, error) {
	data, err := json.Marshal(o)
	if err != nil {
		return nil, err
	}
	return data, err
}

func (o *OpenGraphMetaData) Scan(value interface{}) error {
	if value == nil {
		return errors.New("OG data in database is empty!")
	}
	switch t := value.(type) {
	case string:
		return json.Unmarshal([]byte(t), o)
	case []byte:
		return json.Unmarshal(t, o)
	default:
		return fmt.Errorf("Unknown type for OpenGraphMetaData: %#v", t)
	}
}

type FilterBookmarksAll struct{}

func (f FilterBookmarksAll) BookmarkFilter() {}
