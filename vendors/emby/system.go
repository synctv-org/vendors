package emby

import (
	"net/http"

	json "github.com/json-iterator/go"
)

func (c *Client) SystemInfo() (*SystemInfoResp, error) {
	req, err := c.NewRequest(http.MethodGet, "/emby/System/Info", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var info SystemInfoResp
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}
	return &info, nil
}
