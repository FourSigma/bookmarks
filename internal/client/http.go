package client

import (
	"io"
	"net/http"
	"time"
)

type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
	NewRequest(string, string, io.Reader) (*http.Request, error)
}

func NewHttpClient() HttpClient {
	return &ogHttpClient{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

//Implementation of the HttpClient interface
type ogHttpClient struct {
	client *http.Client
}

func (p ogHttpClient) Do(req *http.Request) (resp *http.Response, err error) {
	return p.client.Do(req)
}

func (p ogHttpClient) NewRequest(method string, path string, body io.Reader) (req *http.Request, err error) {
	if req, err = http.NewRequest(method, path, body); err != nil {
		return
	}
	return
}
