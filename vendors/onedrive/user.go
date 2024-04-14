package onedrive

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

func (c *Client) GetUserInfo() (*MicrosoftUserInfo, error) {
	req, err := c.NewRequest(http.MethodGet, "/me", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	ui := microsoftUserInfoResp{}
	err = json.NewDecoder(resp.Body).Decode(&ui)
	if err != nil {
		return nil, err
	}
	if ui.MicrosoftError.Code != "" {
		return nil, fmt.Errorf("get user info failed: %s", ui.MicrosoftError.Message)
	}
	return &ui.MicrosoftUserInfo, nil
}
