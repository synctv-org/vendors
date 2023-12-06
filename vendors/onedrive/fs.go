package onedrive

import (
	"fmt"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
	"github.com/synctv-org/vendors/utils"
)

func (c *Client) FsGet(path string, opt ...FsOptionFunc) (*FsGetResp, error) {
	opts := FsOption{}
	for _, flo := range opt {
		flo(&opts)
	}
	if path == "" {
		path = "/"
	}
	var (
		req *http.Request
		err error
	)
	if opts.DriveID == "" {
		req, err = c.NewRequest(http.MethodGet, fmt.Sprintf("/me/drive/root:/%s", url.PathEscape(path)), nil)
	} else {
		req, err = c.NewRequest(http.MethodGet, fmt.Sprintf("/drives/%s/root:/%s", opts.DriveID, url.PathEscape(path)), nil)
	}
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var fsg FsGetResp
	err = json.NewDecoder(resp.Body).Decode(&fsg)
	if err != nil {
		return nil, err
	}
	return &fsg, nil
}

type FsOption struct {
	DriveID string
}

type FsOptionFunc func(*FsOption)

func WithDriveID(driveID string) FsOptionFunc {
	return func(o *FsOption) {
		o.DriveID = driveID
	}
}

func (c *Client) FsList(path string, opt ...FsOptionFunc) (*FsListResp, error) {
	opts := FsOption{}
	for _, flo := range opt {
		flo(&opts)
	}
	if path == "" {
		path = "/"
	}
	var (
		req *http.Request
		err error
	)
	if opts.DriveID == "" {
		req, err = c.NewRequest(http.MethodGet, fmt.Sprintf("/me/drive/root:/%s:/children", url.PathEscape(path)), nil)
	} else {
		req, err = c.NewRequest(http.MethodGet, fmt.Sprintf("/drives/%s/root:/%s:/children", opts.DriveID, url.PathEscape(path)), nil)
	}
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var fsl FsListResp
	err = json.NewDecoder(resp.Body).Decode(&fsl)
	if err != nil {
		return nil, err
	}
	return &fsl, nil
}

func (c *Client) FsDownload(path string, opt ...FsOptionFunc) (string, error) {
	opts := FsOption{}
	for _, flo := range opt {
		flo(&opts)
	}
	if path == "" {
		path = "/"
	}
	var (
		req *http.Request
		err error
	)
	if opts.DriveID == "" {
		req, err = c.NewRequest(http.MethodGet, fmt.Sprintf("/me/drive/root:/%s:/content", url.PathEscape(path)), nil)
	} else {
		req, err = c.NewRequest(http.MethodGet, fmt.Sprintf("/drives/%s/root:/%s:/content", opts.DriveID, url.PathEscape(path)), nil)
	}
	if err != nil {
		return "", err
	}
	cli, err := c.HttpClient()
	if err != nil {
		return "", err
	}
	cli = utils.NoRedirectClient(cli)
	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Header.Get("Location"), nil
}
