package onedrive

import (
	"encoding/json"
	"net/http"
)

func (c *Client) DriveList() (*DriveListResp, error) {
	req, err := c.NewRequest(http.MethodGet, "/me/drives", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dl := DriveListResp{}
	err = json.NewDecoder(resp.Body).Decode(&dl)
	if err != nil {
		return nil, err
	}
	return &dl, nil
}

type DriveListResp struct {
	Value []struct {
		ID        string `json:"id"`
		DriveType string `json:"driveType"`
		Name      string `json:"name"`
		Owner     struct {
			User struct {
				ID          string `json:"id"`
				DisplayName string `json:"displayName"`
			} `json:"user"`
		} `json:"owner"`
	} `json:"value"`
}
