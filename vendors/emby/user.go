package emby

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	json "github.com/json-iterator/go"
)

func (c *Client) GetApiKey(username, password string) (string, error) {
	req, err := c.NewRequest(http.MethodPost, "/emby/Users/authenticatebyname", &LoginReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Emby Client=\"Emby Web\", Device=\"SyncTV\", DeviceId=\"%s\", Version=\"4.7.14.0\"", uuid.NewString()))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		return "", errors.New(string(b))
	}
	var loginResp LoginResp
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return "", err
	}
	return loginResp.AccessToken, nil
}

func (c *Client) Login(username, password string) error {
	s, err := c.GetApiKey(username, password)
	if err != nil {
		return err
	}
	c.SetKey(s)
	return nil
}
