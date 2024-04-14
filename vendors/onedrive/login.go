package onedrive

import (
	"golang.org/x/oauth2"
)

func (c *Client) AuthCodeURL(state string) string {
	return c.config.AuthCodeURL(state)
}

func (c *Client) GetToken(code string) (*oauth2.Token, error) {
	return c.config.Exchange(c.ctx, code)
}

func (c *Client) SetToken(tk oauth2.Token) {
	c.tk = tk
	c.tks = c.config.TokenSource(c.ctx, &c.tk)
}

func (c *Client) Login(code string) error {
	tk, err := c.GetToken(code)
	if err != nil {
		return err
	}
	c.SetToken(*tk)
	return nil
}
