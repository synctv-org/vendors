package emby

import (
	"bytes"
	"context"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
	"github.com/zijiren233/go-uhc"
)

type Client struct {
	host   string
	ctx    context.Context
	key    string
	userID string
}

type ClientOption func(*Client)

func WithContext(ctx context.Context) ClientOption {
	return func(c *Client) {
		c.ctx = ctx
	}
}

func WithKey(key string) ClientOption {
	return func(c *Client) {
		c.key = key
	}
}

func WithUserID(userID string) ClientOption {
	return func(c *Client) {
		c.userID = userID
	}
}

func NewClient(host string, opt ...ClientOption) *Client {
	c := Client{
		host: host,
		ctx:  context.Background(),
	}
	for _, o := range opt {
		o(&c)
	}
	return &c
}

func (c *Client) SetKey(key string) {
	c.key = key
}

func (c *Client) SetUserID(userID string) {
	c.userID = userID
}

func (c *Client) NewRequest(method, relative string, data any, querys ...map[string]string) (req *http.Request, err error) {
	result, err := url.JoinPath(c.host, relative)
	if err != nil {
		return nil, err
	}
	if data != nil {
		body := new(bytes.Buffer)
		if err := json.NewEncoder(body).Encode(data); err != nil {
			return nil, err
		}
		req, err = http.NewRequestWithContext(c.ctx, method, result, body)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	} else {
		req, err = http.NewRequestWithContext(c.ctx, method, result, nil)
		if err != nil {
			return nil, err
		}
	}
	if c.key != "" {
		req.Header.Set("X-Emby-Token", c.key)
	}
	q := req.URL.Query()
	for _, query := range querys {
		for k, v := range query {
			q.Add(k, v)
		}
	}
	req.URL.RawQuery = q.Encode()
	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return uhc.Do(req)
}
