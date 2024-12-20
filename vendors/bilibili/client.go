package bilibili

import (
	"context"
	"io"
	"net/http"

	"github.com/synctv-org/vendors/utils"
	"github.com/zijiren233/go-uhc"
)

type Client struct {
	ctx     context.Context
	cookies []*http.Cookie
	buvid   []*http.Cookie
}

type ClientConfig func(*Client)

func WithContext(ctx context.Context) ClientConfig {
	return func(c *Client) {
		c.ctx = ctx
	}
}

func NewClient(cookies []*http.Cookie, conf ...ClientConfig) (*Client, error) {
	cli := &Client{
		cookies: cookies,
		ctx:     context.Background(),
	}
	for _, v := range conf {
		v(cli)
	}
	b, err := getBuvidCookies(cli.ctx)
	if err != nil {
		return nil, err
	}
	cli.buvid = b
	return cli, nil
}

func (c *Client) SetCookies(cookies []*http.Cookie) {
	c.cookies = cookies
}

type RequestConfig struct {
	wbi bool
}

func defaultRequestConfig() *RequestConfig {
	return &RequestConfig{
		wbi: true,
	}
}

type RequestOption func(*RequestConfig)

func WithoutWbi() RequestOption {
	return func(c *RequestConfig) {
		c.wbi = false
	}
}

func (c *Client) NewRequest(method, url string, body io.Reader, conf ...RequestOption) (req *http.Request, err error) {
	config := defaultRequestConfig()
	for _, v := range conf {
		v(config)
	}
	if config.wbi {
		url, err = signAndGenerateURL(c.ctx, url)
		if err != nil {
			return nil, err
		}
	}
	req, err = http.NewRequestWithContext(c.ctx, method, url, body)
	if err != nil {
		return nil, err
	}
	for _, cookie := range c.buvid {
		req.AddCookie(cookie)
	}
	for _, cookie := range c.cookies {
		req.AddCookie(cookie)
	}
	req.Header.Set("User-Agent", utils.UA)
	req.Header.Set("Referer", "https://www.bilibili.com")
	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return uhc.Do(req)
}
