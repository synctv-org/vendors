package bilibili

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	json "github.com/json-iterator/go"
)

// https://api.live.bilibili.com/room/v1/Room/get_info
//
// room_id: room id, can be sort id
func (c *Client) ParseLivePage(roomID uint64) (*VideoPageInfo, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.live.bilibili.com",
		Path:   "/room/v1/Room/get_info",
	}
	q := url.Values{}
	q.Set("room_id", strconv.FormatUint(roomID, 10))
	u.RawQuery = q.Encode()

	req, err := c.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data parseLivePageResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, fmt.Errorf("parse live page failed: %s, code: %d", data.Message, data.Code)
	}

	uinfo, err := c.GetLiveMasterInfo(data.Data.Uid)
	if err != nil {
		return nil, fmt.Errorf("get live master info failed: %w", err)
	}

	return &VideoPageInfo{
		Title:  data.Data.Title,
		Actors: uinfo.Username,
		VideoInfos: []*VideoInfo{
			{
				Cid:        data.Data.RoomID,
				Name:       data.Data.Title,
				CoverImage: data.Data.UserCover,
				Live:       true,
			},
		},
	}, nil
}

type LiveMasterInfo struct {
	UID      uint64 `json:"uid"`
	Username string `json:"username"`
	Face     string `json:"face"`
}

// https://api.live.bilibili.com/live_user/v1/Master/info
func (c *Client) GetLiveMasterInfo(uid uint64) (*LiveMasterInfo, error) {
	u := url.URL{
		Scheme: "https",
		Host:   "api.live.bilibili.com",
		Path:   "/live_user/v1/Master/info",
	}
	q := url.Values{}
	q.Set("uid", strconv.FormatUint(uid, 10))
	u.RawQuery = q.Encode()

	req, err := c.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data getLiveMasterInfoResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, fmt.Errorf("get live master info failed: %s, code: %d", data.Message, data.Code)
	}

	return &LiveMasterInfo{
		UID:      data.Data.Info.Uid,
		Username: data.Data.Info.Username,
		Face:     data.Data.Info.Face,
	}, nil
}

type LiveStream struct {
	Quality uint64   `json:"quality"`
	Urls    []string `json:"url"`
	Desc    string   `json:"desc"`
}

type GetLiveStreamWithQualityResp struct {
	AcceptQuality []uint64 `json:"acceptQuality"`
	*LiveStream
}

// https://api.live.bilibili.com/room/v1/Room/playUrl?cid=1000&quality=4&platform=web
//
// cid: room id, not sort id
//
// platform:
//
//	web: http-flv
//	h5: hls
//
// qn:
//
//	2: 流畅
//	3: 高清
//	4: 原画
func (c *Client) GetLiveStreamWithQuality(cid uint64, hls bool, qn uint64) (*GetLiveStreamWithQualityResp, error) {
	var platform string
	if hls {
		platform = "h5"
	} else {
		platform = "web"
	}

	u := url.URL{
		Scheme: "https",
		Host:   "api.live.bilibili.com",
		Path:   "/room/v1/Room/playUrl",
	}
	q := url.Values{}
	q.Set("cid", strconv.FormatUint(cid, 10))
	q.Set("quality", strconv.FormatUint(qn, 10))
	q.Set("platform", platform)
	u.RawQuery = q.Encode()

	req, err := c.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data getLiveStreamResp
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if data.Code != 0 {
		return nil, fmt.Errorf("get live stream failed: %s, code: %d", data.Message, data.Code)
	}

	var desc string
	for _, v := range data.Data.QualityDescription {
		if v.Quality == data.Data.CurrentQn {
			desc = v.Desc
			break
		}
	}

	urls := make([]string, len(data.Data.DUrl))
	for i, v := range data.Data.DUrl {
		urls[i] = v.Url
	}

	var acceptQuality []uint64 = make([]uint64, len(data.Data.AcceptQuality))
	for i, v := range data.Data.AcceptQuality {
		q, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse accept quality failed: %w", err)
		}
		acceptQuality[i] = q
	}

	return &GetLiveStreamWithQualityResp{
		AcceptQuality: acceptQuality,
		LiveStream: &LiveStream{
			Quality: data.Data.CurrentQuality,
			Urls:    urls,
			Desc:    desc,
		},
	}, nil
}

// https://api.live.bilibili.com/room/v1/Room/playUrl?cid=1000&quality=4&platform=web
//
// cid: room id, not sort id
//
// platform:
//
//	web: http-flv
//	h5: hls
func (c *Client) GetLiveStreams(cid uint64, hls bool) ([]*LiveStream, error) {
	first, err := c.GetLiveStreamWithQuality(cid, hls, 0)
	if err != nil {
		return nil, err
	}
	streams := make([]*LiveStream, 0, len(first.AcceptQuality))
	streams = append(streams, first.LiveStream)

	for _, qn := range first.AcceptQuality {
		if qn == first.LiveStream.Quality {
			continue
		}

		stream, err := c.GetLiveStreamWithQuality(cid, hls, qn)
		if err != nil {
			return nil, err
		}
		streams = append(streams, stream.LiveStream)
	}

	return streams, nil
}