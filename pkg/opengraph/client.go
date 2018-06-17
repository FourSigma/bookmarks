package opengraph

import (
	"context"
	"fmt"
	"net/http"
	neturl "net/url"
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

type ErrorOgClient struct {
	Code    int
	Message string
}

func (e ErrorOgClient) Error() string {
	return fmt.Sprintf("URL returned with status code %s (%s)", e.Code, e.Message)
}

func NewOgClientError(code int, msg string) ErrorOgClient {
	return ErrorOgClient{
		Code:    code,
		Message: msg,
	}
}

//Implementation of the HttpClient interface
type ogClient struct {
	client *http.Client
}

func (p ogClient) OpenGraphMetaData(ctx context.Context, url string) (data map[string]string, err error) {
	if _, err = neturl.Parse(url); err != nil {
		return nil, NewOgClientError(http.StatusBadRequest, "Malformed URL")
	}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, NewOgClientError(http.StatusInternalServerError, "Internal Server Error")
	}
	if resp.StatusCode >= 300 {
		return nil, NewOgClientError(resp.StatusCode, resp.Status)
	}
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
