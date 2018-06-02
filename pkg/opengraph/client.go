package opengraph

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/johnreutersward/opengraph"
)

type OGClient interface {
	OpenGraphMetaData(context.Context, string) (map[string]string, error)
}

func NewOGClient() OGClient {
	return &ogClient{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
	}
}

//Implementation of the HttpClient interface
type ogClient struct {
	client *http.Client
}

func (p ogClient) OpenGraphMetaData(ctx context.Context, url string) (data map[string]string, err error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, err
	}
	fmt.Println(resp.StatusCode, resp.Status)
	defer resp.Body.Close()

	md, err := opengraph.Extract(resp.Body)
	if err != nil {
		return nil, err
	}

	data = map[string]string{}
	for _, v := range md {
		fmt.Println(v.Property, ": ", v.Content)
		data[v.Property] = v.Content
	}

	return
}
