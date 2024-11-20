package bilibili

import (
	"errors"
	"net/http"

	json "github.com/json-iterator/go"
)

type UserInfo struct {
	Username string `json:"username"`
	Face     string `json:"face"`
	IsLogin  bool   `json:"isLogin"`
	IsVip    bool   `json:"isVip"`
}

func (c *Client) UserInfo() (*UserInfo, error) {
	req, err := c.NewRequest(http.MethodGet, "https://api.bilibili.com/x/web-interface/nav", nil, WithoutWbi())
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	var nav Nav
	err = json.NewDecoder(resp.Body).Decode(&nav)
	if err != nil {
		return nil, err
	}
	if nav.Code != 0 {
		return nil, errors.New(nav.Message)
	}
	return &UserInfo{
		IsLogin:  nav.Data.IsLogin,
		Username: nav.Data.Uname,
		Face:     nav.Data.Face,
		IsVip:    nav.Data.VipStatus == 1,
	}, nil
}
