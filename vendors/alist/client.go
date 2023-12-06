package alist

import (
	"bytes"
	"context"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
)

type Client struct {
	host       string
	token      string
	httpClient *http.Client
	ctx        context.Context
}

type ClientConfig func(*Client)

func WithHttpClient(httpClient *http.Client) ClientConfig {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}

func WithContext(ctx context.Context) ClientConfig {
	return func(c *Client) {
		c.ctx = ctx
	}
}

func NewClient(host, token string, conf ...ClientConfig) (*Client, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	cli := &Client{
		host:       u.String(),
		token:      token,
		httpClient: http.DefaultClient,
		ctx:        context.Background(),
	}
	for _, v := range conf {
		v(cli)
	}
	return cli, nil
}

func (c *Client) SetToken(token string) {
	c.token = token
}

func (c *Client) NewRequest(method, relative string, data any) (req *http.Request, err error) {
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
		req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	} else {
		req, err = http.NewRequestWithContext(c.ctx, method, result, nil)
		if err != nil {
			return nil, err
		}
	}

	if c.token != "" {
		req.Header.Set("Authorization", c.token)
	}
	req.Header.Set("Origin", c.host)
	req.Header.Set("Referer", c.host)
	return req, nil
}
