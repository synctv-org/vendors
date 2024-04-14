package onedrive

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) DriveList() (*DriveList, error) {
	req, err := c.NewRequest(http.MethodGet, "/me/drives", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	dl := driveListResp{}
	err = json.NewDecoder(resp.Body).Decode(&dl)
	if err != nil {
		return nil, err
	}
	if dl.MicrosoftError.Code != "" {
		return nil, fmt.Errorf("get drive list failed: %s", dl.MicrosoftError.Message)
	}
	return &dl.DriveList, nil
}

type DriveList struct {
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

type driveListResp struct {
	DriveList
	MicrosoftError `json:"error"`
}
