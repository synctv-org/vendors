package onedrive

import (
	"fmt"
	"net/http"
	"net/url"

	json "github.com/json-iterator/go"
	"github.com/synctv-org/vendors/utils"
)

func (c *Client) FsGet(path string, opt ...FsOptionFunc) (*FsGet, error) {
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
	var fsg fsGetResp
	err = json.NewDecoder(resp.Body).Decode(&fsg)
	if err != nil {
		return nil, err
	}
	if fsg.MicrosoftError.Code != "" {
		return nil, fmt.Errorf("get file info failed: %s", fsg.MicrosoftError.Message)
	}
	return &fsg.FsGet, nil
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

func (c *Client) FsList(path string, opt ...FsOptionFunc) (*FsList, error) {
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
	var fsl fsListResp
	err = json.NewDecoder(resp.Body).Decode(&fsl)
	if err != nil {
		return nil, err
	}
	if fsl.MicrosoftError.Code != "" {
		return nil, fmt.Errorf("get file list failed: %s", fsl.MicrosoftError.Message)
	}
	return &fsl.FsList, nil
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
	cli := utils.NoRedirectClient(c.httpClient())
	resp, err := cli.Do(req)
	if err != nil {
		return "", err
	}
	resp.Body.Close()
	location := resp.Header.Get("Location")
	if location == "" {
		return "", fmt.Errorf("download file failed: location not found")
	}
	return location, nil
}
