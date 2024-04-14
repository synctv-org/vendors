package alist

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

func (c *Client) Me() (*meResp, error) {
	req, err := c.NewRequest(http.MethodGet, "/api/me", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r MeResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("alist: %s", r.Message)
	}
	return &r.Data, nil
}
