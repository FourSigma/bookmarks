package core

import (
	"context"
	"net/url"

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
	Id   uuid.UUID
	URL  url.URL //Must be unique
	Hash string
	Data struct {
		Title       string `json:"title"`
		Type        string `json:"type"`
		Url         string `json:"url"`
		Site        string `json:"site"`
		SiteName    string `json:"site_name,omitempty"`
		Description string `json:"description"`
		Locale      string `json:"locale"`
		Image       string `json:"image,omitempty"`
		Content     string `json:"content"`
	}
}
