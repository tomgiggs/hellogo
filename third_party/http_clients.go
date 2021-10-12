package third_party

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/parnurzeal/gorequest"
	"net/http"
	"time"
)

const (
	RetryCounts             = 2
	RetryInterval           = 3 * time.Second
	componentIDGOHttpClient = 5005
)

type AdvanceClient struct {
}

func (ac AdvanceClient) Get() {
	request := gorequest.New()
	resp, body, errs := request.Get("http://example.com/").End()
	fmt.Println(resp, body, errs)
}

type ClientConfig struct {
	ctx       context.Context
	client    *resty.Client
	extraTags map[string]string
}

type ClientOption func(*ClientConfig)

func WithClientTimeout(ttl int64) ClientOption {
	return func(c *ClientConfig) {
		if ttl > 0 {
			c.client.SetTimeout(time.Duration(ttl) * time.Second)
		}
	}
}

func WithClient(client *resty.Client) ClientOption {
	return func(c *ClientConfig) {
		c.client = client
	}
}

func WithContext(ctx context.Context) ClientOption {
	return func(c *ClientConfig) {
		c.ctx = ctx
	}
}

type transport struct {
	*ClientConfig
	delegated http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	resp, err = t.delegated.RoundTrip(req)
	if err != nil {
		return
	}
	return resp, nil
}

func NewRestyClient(ctx context.Context, options ...ClientOption) (client *resty.Client) {
	client = resty.New()
	client.SetDebug(true).EnableTrace()
	co := &ClientConfig{ctx: ctx}
	for _, option := range options {
		option(co)
	}
	if co.client == nil {
		co.client = client
	}
	tp := &transport{
		ClientConfig: co,
		delegated:    http.DefaultTransport,
	}
	if co.client.GetClient().Transport != nil {
		tp.delegated = co.client.GetClient().Transport
	}
	co.client.SetTransport(tp)

	client.OnBeforeRequest(func(ct *resty.Client, req *resty.Request) error {
		return nil // if its success otherwise return error
	})
	client.OnAfterResponse(func(ct *resty.Client, resp *resty.Response) error {
		return nil
	})
	return client
}


