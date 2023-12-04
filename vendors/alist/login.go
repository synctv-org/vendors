package alist

import (
	"context"
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

func Login(ctx context.Context, host string, data LoginReq) (string, error) {
	cli, err := NewClient(host, "", WithContext(ctx))
	if err != nil {
		return "", err
	}
	req, err := cli.NewRequest(http.MethodPost, "/api/auth/login", &data)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var r LoginResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return "", err
	}
	if r.Code != 200 {
		return "", fmt.Errorf("alist: %s", r.Message)
	}
	return r.Data.Token, nil
}

func (c *Client) Login(data LoginReq) error {
	s, err := Login(c.ctx, c.host, data)
	if err != nil {
		return err
	}
	c.SetToken(s)
	return nil
}
