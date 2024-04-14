package alist

import (
	"fmt"
	"net/http"

	json "github.com/json-iterator/go"
)

func (c *Client) FsGet(fSGetReq *FsGetReq) (*fsGetResp, error) {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/get", fSGetReq)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r FsGetResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("alist: %s", r.Message)
	}
	return &r.Data, nil
}

func (c *Client) FsList(fSListReq *FsListReq) (*fsListResp, error) {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/list", fSListReq)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r FsListResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("alist: %s", r.Message)
	}
	return &r.Data, nil
}

func (c *Client) FsOther(fSOtherReq *FsOtherReq) (*fsOtherResp, error) {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/other", fSOtherReq)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r FsOtherResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("alist: %s", r.Message)
	}
	return &r.Data, nil
}

func (c *Client) FsMkdir(fSMkdirReq *FsMkdirReq) error {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/mkdir", fSMkdirReq)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var r FsMkdirResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	if r.Code != 200 {
		return fmt.Errorf("alist: %s", r.Message)
	}
	return nil
}

func (c *Client) FsRename(fSRenameReq *FsRenameReq) error {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/rename", fSRenameReq)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var r FsRenameResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	if r.Code != 200 {
		return fmt.Errorf("alist: %s", r.Message)
	}
	return nil
}

func (c *Client) FsRemove(fSRemoveReq *FsRemoveReq) error {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/remove", fSRemoveReq)
	if err != nil {
		return err
	}
	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var r FsRemoveResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return err
	}
	if r.Code != 200 {
		return fmt.Errorf("alist: %s", r.Message)
	}
	return nil
}

func (c *Client) FsSearch(fSSearchReq *FsSearchReq) (*fsSearchResp, error) {
	req, err := c.NewRequest(http.MethodPost, "/api/fs/search", fSSearchReq)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var r FsSearchResp
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, err
	}
	if r.Code != 200 {
		return nil, fmt.Errorf("alist: %s", r.Message)
	}
	return &r.Data, nil
}
