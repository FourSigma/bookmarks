package core

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type LinkService interface {
	Create(context.Context, *Link) error
	Get(context.Context, LinkId) (Link, error)
	List(context.Context, LinkFilter, ...Opt) ([]Link, error)
	Update(context.Context, LinkId, *Link) error
	Delete(context.Context, LinkId) error

	GetLinkFromURL(context.Context, string) (Link, error)
}

type LinkFilter interface {
	LinkFilter()
}

type Opt interface {
	Opt()
}

type LinkId uuid.UUID
type Link struct {
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

func (l Link) Valid() error {
	return nil
}
