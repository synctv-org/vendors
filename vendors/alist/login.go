package alist

import (
	"context"
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

type LoginOptions struct {
	Hashed bool
}

type LoginOpt func(*LoginOptions)

func WithHashed() LoginOpt {
	return func(o *LoginOptions) {
		o.Hashed = true
	}
}

func Login(ctx context.Context, host string, data LoginReq, opts ...LoginOpt) (string, error) {
	var opt LoginOptions
	for _, o := range opts {
		o(&opt)
	}
	cli, err := NewClient(host, "", WithContext(ctx))
	if err != nil {
		return "", err
	}
	var req *http.Request
	if opt.Hashed {
		req, err = cli.NewRequest(http.MethodPost, "/api/auth/login/hash", &data)
	} else {
		req, err = cli.NewRequest(http.MethodPost, "/api/auth/login", &data)
	}
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

func (c *Client) Login(data LoginReq, opts ...LoginOpt) error {
	s, err := Login(c.ctx, c.host, data, opts...)
	if err != nil {
		return err
	}
	c.SetToken(s)
	return nil
}
