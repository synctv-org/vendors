package emby

import (
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
	json "github.com/json-iterator/go"
)

type GetApiKeyResp struct {
	AccessToken string `json:"AccessToken"`
	UserID      string `json:"User"`
	ServerID    string `json:"ServerId"`
}

func (c *Client) GetApiKey(username, password string) (*GetApiKeyResp, error) {
	req, err := c.NewRequest(http.MethodPost, "/emby/Users/authenticatebyname", &LoginReq{
		Username: username,
		Password: password,
	})
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Emby Client=\"Emby Web\", Device=\"SyncTV\", DeviceId=\"%s\", Version=\"4.7.14.0\"", uuid.NewString()))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}
	var loginResp LoginResp
	if err := json.NewDecoder(resp.Body).Decode(&loginResp); err != nil {
		return nil, err
	}
	return &GetApiKeyResp{
		AccessToken: loginResp.AccessToken,
		UserID:      loginResp.User.ID,
		ServerID:    loginResp.ServerID,
	}, nil
}

func (c *Client) Login(username, password string) error {
	s, err := c.GetApiKey(username, password)
	if err != nil {
		return err
	}
	c.SetKey(s.AccessToken)
	c.SetUserID(s.UserID)
	return nil
}

func (c *Client) Logout() error {
	req, err := c.NewRequest(http.MethodPost, "/emby/Sessions/Logout", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Emby Client=\"Emby Web\", Device=\"SyncTV\", DeviceId=\"%s\", Version=\"4.7.14.0\"", uuid.NewString()))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}
	c.SetKey("")
	c.SetUserID("")
	return nil
}

func (c *Client) Me() (*MeResp, error) {
	if c.userID == "" {
		return nil, errors.New("user id not set")
	}
	req, err := c.NewRequest(http.MethodGet, fmt.Sprintf("/emby/Users/%s", c.userID), nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("status code %d: %s", resp.StatusCode, string(b))
	}
	var meResp MeResp
	if err := json.NewDecoder(resp.Body).Decode(&meResp); err != nil {
		return nil, err
	}
	return &meResp, nil
}
