package emby

import (
	"fmt"
	"io"
	"net/http"

	json "github.com/json-iterator/go"
)

func (c *Client) GetItem(id string) (*MediaItem, error) {
	if id == "" || id == "0" {
		id = "1"
	}
	req, err := c.NewRequest(http.MethodGet, "/emby/Items", nil, map[string]string{
		"Ids":    id,
		"Fields": "MediaSources,ParentId",
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

// visual path:
// CollectionFolder:
//
//	1/{Id}
//	Home/{Name}
//
// Series:
//
//	1/{ParentId}/{Id}
//	Home/{GetItem(ParentId).Name}/{Name}
//
// Season:
//
//	1/{GetItem(SeriesId).ParentId}/{SeriesId}/{Id}
//	Home/{GetItem(GetItem(SeriesId).ParentId).Name}/{SeriesName}/{Name}
//
// Episode:
//
//	1/{GetItem(SeriesId).ParentId}/{SeriesId}/{SeasonId}/{Id}
//	Home/{GetItem(GetItem(SeriesId).ParentId).Name}/{SeriesName}/{SeasonName}/{Name}
func (c *Client) GetItems(id string) (*GetItemResp, error) {
	if id == "" || id == "0" {
		id = "1"
	}
	req, err := c.NewRequest(http.MethodGet, "/emby/Items", nil, map[string]string{
		"ParentId": id,
		"Fields":   "MediaSources,ParentId",
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
