package emby

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	json "github.com/json-iterator/go"
)

func (c *Client) GetItem(id string) (*Items, error) {
	if id == "" || id == "0" {
		id = "1"
	}
	req, err := c.NewRequest(http.MethodGet, "/emby/Items", nil, map[string]string{
		"Ids":    id,
		"Fields": "MediaSources,ParentId,Container",
	})
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
	var fsGetResp GetItemResp
	if err := json.NewDecoder(resp.Body).Decode(&fsGetResp); err != nil {
		return nil, err
	}
	if len(fsGetResp.Items) == 0 {
		return nil, fmt.Errorf("item not found")
	} else if len(fsGetResp.Items) > 1 {
		return nil, fmt.Errorf("item not unique")
	}
	return &fsGetResp.Items[0], nil
}

type GetItemsOptionFunc func(map[string]string)

func WithStartIndex(startIndex uint64) GetItemsOptionFunc {
	return func(o map[string]string) {
		o["StartIndex"] = fmt.Sprintf("%d", startIndex)
	}
}

func WithLimit(limit uint64) GetItemsOptionFunc {
	return func(o map[string]string) {
		o["Limit"] = fmt.Sprintf("%d", limit)
	}
}

func WithParentId(parentId string) GetItemsOptionFunc {
	return func(o map[string]string) {
		if parentId == "" || parentId == "0" {
			parentId = "1"
		}
		o["ParentId"] = parentId
	}
}

func WithSortBy(sortBy string) GetItemsOptionFunc {
	return func(o map[string]string) {
		o["SortBy"] = sortBy
	}
}

func WithSortOrderAsc() GetItemsOptionFunc {
	return func(o map[string]string) {
		o["SortOrder"] = "Ascending"
	}
}

func WithSortOrderDesc() GetItemsOptionFunc {
	return func(o map[string]string) {
		o["SortOrder"] = "Descending"
	}
}

func WithSearch(search string) GetItemsOptionFunc {
	return func(o map[string]string) {
		o["SearchTerm"] = search
	}
}

// 递归获取
func WithRecursive() GetItemsOptionFunc {
	return func(o map[string]string) {
		o["Recursive"] = "true"
	}
}

func WithIncludeItemTypes(types ...string) GetItemsOptionFunc {
	return func(o map[string]string) {
		if v, ok := o["IncludeItemTypes"]; ok {
			o["IncludeItemTypes"] = fmt.Sprintf("%s,%s", v, strings.Join(types, ","))
		} else {
			o["IncludeItemTypes"] = strings.Join(types, ",")
		}
	}
}

func (c *Client) GetItems(opt ...GetItemsOptionFunc) (*GetItemResp, error) {
	o := map[string]string{
		"Fields": "MediaSources,ParentId,Container",
	}
	for _, f := range opt {
		f(o)
	}
	if _, ok := o["SearchTerm"]; ok {
		if i, ok := o["ParentId"]; ok && i == "1" {
			delete(o, "ParentId")
		}
	}
	req, err := c.NewRequest(http.MethodGet, "/emby/Items", nil, o)
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
	var fsGetResp GetItemResp
	if err := json.NewDecoder(resp.Body).Decode(&fsGetResp); err != nil {
		return nil, err
	}
	return &fsGetResp, nil
}

func (c *Client) Library() (*LibraryResp, error) {
	req, err := c.NewRequest(http.MethodGet, "/emby/Library/MediaFolders", nil)
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
	var fsListResp LibraryResp
	if err := json.NewDecoder(resp.Body).Decode(&fsListResp); err != nil {
		return nil, err
	}
	return &fsListResp, nil
}
