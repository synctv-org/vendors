package onedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/synctv-org/vendors/utils"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"
)

func newOauth2Config(clientID, clientSecret, redirectURL string) oauth2.Config {
	c := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"user.read", "files.read", "offline_access"},
		Endpoint:     microsoft.LiveConnectEndpoint,
	}
	return c
}

type Client struct {
	config  oauth2.Config
	ctx     context.Context
	tk      oauth2.Token
	tks     oauth2.TokenSource
	baseURL string
}

type ClientOption func(*Client)

func WithContext(ctx context.Context) ClientOption {
	return func(c *Client) {
		c.ctx = ctx
	}
}

func WithToken(tk oauth2.Token) ClientOption {
	return func(c *Client) {
		c.SetToken(tk)
	}
}

func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

func NewClient(clientID, clientSecret, redirectURL string, opt ...ClientOption) *Client {
	c := Client{
		config:  newOauth2Config(clientID, clientSecret, redirectURL),
		ctx:     context.Background(),
		baseURL: "https://graph.microsoft.com/v1.0",
	}
	for _, o := range opt {
		o(&c)
	}
	return &c
}

func (c *Client) Token() (*oauth2.Token, error) {
	return c.tks.Token()
}

func (c *Client) httpClient() *http.Client {
	return &http.Client{
		Transport: &oauth2.Transport{
			Source: c.tks,
			Base:   utils.DefaultUtlsHttpRoundTripper,
		},
	}
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient().Do(req)
}

func (c *Client) NewRequest(method, relative string, data any) (req *http.Request, err error) {
	result, err := url.JoinPath(c.baseURL, relative)
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
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequestWithContext(c.ctx, method, result, nil)
		if err != nil {
			return nil, err
		}
	}

	return req, nil
}
