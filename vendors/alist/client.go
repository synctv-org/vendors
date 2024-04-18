package alist

import (
	"bytes"
	"context"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
	"github.com/synctv-org/vendors/utils"
	"github.com/zijiren233/go-uhc"
)

type Client struct {
	host      string
	token     string
	ctx       context.Context
	userAgent string
}

type ClientConfig func(*Client)

func WithContext(ctx context.Context) ClientConfig {
	return func(c *Client) {
		c.ctx = ctx
	}
}

func WithUserAgent(userAgent string) ClientConfig {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

func NewClient(host, token string, conf ...ClientConfig) (*Client, error) {
	u, err := url.Parse(host)
	if err != nil {
		return nil, err
	}
	cli := &Client{
		host:  u.String(),
		token: token,
		ctx:   context.Background(),
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
	if req.Header.Get("User-Agent") == "" {
		if c.userAgent != "" {
			req.Header.Set("User-Agent", c.userAgent)
		} else {
			req.Header.Set("User-Agent", utils.UA)
		}
	}
	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return uhc.Do(req)
}
