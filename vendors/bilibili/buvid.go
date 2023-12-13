package bilibili

import (
	"context"
	"errors"
	"net/http"
	"time"

	json "github.com/json-iterator/go"
	"github.com/synctv-org/vendors/utils"
	"github.com/zijiren233/gencontainer/refreshcache"
)

type buvid struct {
	b3, b4 string
}

var buvidCache = refreshcache.NewRefreshCache[*buvid](func(ctx context.Context, args ...struct{}) (*buvid, error) {
	b3, b4, err := newBuvid(ctx)
	if err != nil {
		return nil, err
	}
	return &buvid{
		b3: b3,
		b4: b4,
	}, nil
}, time.Hour)

func getBuvidCookies(ctx context.Context) ([]*http.Cookie, error) {
	buvid, err := buvidCache.Get(ctx)
	if err != nil {
		return nil, err
	}
	return []*http.Cookie{
		{
			Name:  "buvid3",
			Value: buvid.b3,
		},
		{
			Name:  "buvid4",
			Value: buvid.b4,
		},
	}, nil
}

type spiResp struct {
	Code int `json:"code"`
	Data struct {
		B3 string `json:"b_3"`
		B4 string `json:"b_4"`
	} `json:"data"`
	Message string `json:"message"`
}

func newBuvid(ctx context.Context) (string, string, error) {
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://api.bilibili.com/x/frontend/finger/spi", nil)
	if err != nil {
		return "", "", err
	}
	r.Header.Set("User-Agent", utils.UA)
	r.Header.Set("Referer", "https://www.bilibili.com")
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()
	var b spiResp
	err = json.NewDecoder(resp.Body).Decode(&b)
	if err != nil {
		return "", "", err
	}
	if b.Code != 0 {
		return "", "", errors.New(b.Message)
	}
	return b.Data.B3, b.Data.B4, nil
}
