package onedrive

import (
	"net/http"

	json "github.com/json-iterator/go"
)

func (c *Client) GetUserInfo() (*microsoftUserInfo, error) {
	req, err := c.NewRequest(http.MethodGet, "/me", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ui := microsoftUserInfo{}
	err = json.NewDecoder(resp.Body).Decode(&ui)
	if err != nil {
		return nil, err
	}
	return &ui, nil
}
