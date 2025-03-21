package bilibili

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/zencoder/go-dash/v3/mpd"
)

type VideoPageInfo struct {
	Title      string       `json:"title"`
	Actors     string       `json:"actors"`
	VideoInfos []*VideoInfo `json:"videoInfos"`
}

type VideoInfo struct {
	Bvid       string `json:"bvid,omitempty"`
	Name       string `json:"name"`
	CoverImage string `json:"coverImage"`
	Cid        uint64 `json:"cid,omitempty"`
	Epid       uint64 `json:"epid,omitempty"`
	Live       bool   `json:"live"`
}

type ParseVideoPageConf struct {
	GetSections bool
}

type ParseVideoPageConfig func(*ParseVideoPageConf)

func WithGetSections(getSections bool) ParseVideoPageConfig {
	return func(c *ParseVideoPageConf) {
		c.GetSections = getSections
	}
}

func (c *Client) ParseVideoPage(aid uint64, bvid string, conf ...ParseVideoPageConfig) (*VideoPageInfo, error) {
	config := &ParseVideoPageConf{}
	for _, v := range conf {
		v(config)
	}
	var url string
	if aid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?aid=%d", aid)
	} else if bvid != "" {
		url = "https://api.bilibili.com/x/web-interface/view?bvid=" + bvid
	} else {
		return nil, errors.New("aid and bvid are both empty")
	}
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := videoPageInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(info.Message)
	}
	r := &VideoPageInfo{
		Title:  info.Data.Title,
		Actors: info.Data.Owner.Name,
	}

	if config.GetSections && len(info.Data.UgcSeason.Sections) != 0 {
		r.Title = info.Data.UgcSeason.Title
		for _, v := range info.Data.UgcSeason.Sections {
			for _, episode := range v.Episodes {
				r.VideoInfos = append(r.VideoInfos, &VideoInfo{
					Bvid:       episode.Bvid,
					Cid:        episode.Cid,
					Name:       episode.Title,
					CoverImage: episode.Arc.Pic,
				})
			}
		}
	} else {
		r.VideoInfos = make([]*VideoInfo, len(info.Data.Pages))
		if len(info.Data.Pages) == 1 {
			info.Data.Pages[0].Part = info.Data.Title
		}
		for i, page := range info.Data.Pages {
			r.VideoInfos[i] = &VideoInfo{
				Bvid:       info.Data.Bvid,
				Cid:        page.Cid,
				Name:       page.Part,
				CoverImage: info.Data.Pic,
			}
		}
	}
	return r, nil
}

const (
	Q240P    uint64 = 6
	Q360P    uint64 = 16
	Q480P    uint64 = 32
	Q720P    uint64 = 64
	Q1080P   uint64 = 80
	Q1080PP  uint64 = 112
	Q1080P60 uint64 = 116
	Q4K      uint64 = 120
	QHDR     uint64 = 124
	QDOLBY   uint64 = 126
	Q8K      uint64 = 127
)

type VideoURL struct {
	URL               string   `json:"url"`
	AcceptDescription []string `json:"acceptDescription"`
	AcceptQuality     []uint64 `json:"acceptQuality"`
	CurrentQuality    uint64   `json:"currentQuality"`
}

type GetVideoURLConf struct {
	Quality uint64
}

func (c *GetVideoURLConf) fix() {
	if c.Quality == 0 {
		c.Quality = Q1080PP
	}
}

type GetVideoURLConfig func(*GetVideoURLConf)

func WithQuality(q uint64) GetVideoURLConfig {
	return func(c *GetVideoURLConf) {
		c.Quality = q
	}
}

// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/video/videostream_url.md
func (c *Client) GetVideoURL(aid uint64, bvid string, cid uint64, conf ...GetVideoURLConfig) (*VideoURL, error) {
	config := &GetVideoURLConf{
		Quality: Q1080PP,
	}
	for _, v := range conf {
		v(config)
	}
	config.fix()

	var url string
	if aid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/x/player/wbi/playurl?aid=%d&cid=%d&qn=%d&platform=html5&high_quality=1", aid, cid, config.Quality)
	} else if bvid != "" {
		url = fmt.Sprintf("https://api.bilibili.com/x/player/wbi/playurl?bvid=%s&cid=%d&qn=%d&platform=html5&high_quality=1", bvid, cid, config.Quality)
	} else {
		return nil, errors.New("aid and bvid are both empty")
	}
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := videoInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(info.Message)
	}
	return &VideoURL{
		AcceptDescription: info.Data.AcceptDescription,
		AcceptQuality:     info.Data.AcceptQuality,
		CurrentQuality:    info.Data.Quality,
		URL:               info.Data.Durl[0].URL,
	}, nil
}

type GetDashVideoURLConf struct {
	HDR bool
}

type GetDashVideoURLConfig func(*GetDashVideoURLConf)

func WithHDR(hdr bool) GetDashVideoURLConfig {
	return func(c *GetDashVideoURLConf) {
		c.HDR = hdr
	}
}

// https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/video/videostream_url.md
func (c *Client) GetDashVideoURL(aid uint64, bvid string, cid uint64, conf ...GetDashVideoURLConfig) (*mpd.MPD, *mpd.MPD, error) {
	config := &GetDashVideoURLConf{}
	for _, v := range conf {
		v(config)
	}

	var (
		fnval    uint64 = 16
		extQuery string
	)

	if config.HDR {
		fnval |= 64
	}

	// 4K
	fnval |= 128
	// 8K
	fnval |= 1024

	var url string
	if aid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/x/player/wbi/playurl?aid=%d&cid=%d&fnver=0&platform=pc&fnval=%d%s", aid, cid, fnval, extQuery)
	} else if bvid != "" {
		url = fmt.Sprintf("https://api.bilibili.com/x/player/wbi/playurl?bvid=%s&cid=%d&fnver=0&platform=pc&fnval=%d%s", bvid, cid, fnval, extQuery)
	} else {
		return nil, nil, errors.New("aid and bvid are both empty")
	}
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	info := dashResp{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, nil, err
	}
	if info.Code != 0 {
		return nil, nil, errors.New(info.Message)
	}
	videoFormatsID2Name := map[int]string{}
	for _, v := range info.Data.SupportFormats {
		videoFormatsID2Name[v.Quality] = v.NewDescription
	}
	m := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, fmt.Sprintf("PT%.2fS", info.Data.Dash.Duration), fmt.Sprintf("PT%.2fS", info.Data.Dash.MinBufferTime))
	hevcM := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, fmt.Sprintf("PT%.2fS", info.Data.Dash.Duration), fmt.Sprintf("PT%.2fS", info.Data.Dash.MinBufferTime))
	var as, mAs, hevcMAs *mpd.AdaptationSet
	for _, v := range info.Data.Dash.Video {
		if as == nil {
			mAs, err = m.AddNewAdaptationSetVideo(v.MimeType, "progressive", true, v.StartWithSap)
			if err != nil {
				return nil, nil, err
			}
			hevcMAs, err = hevcM.AddNewAdaptationSetVideo(v.MimeType, "progressive", true, v.StartWithSap)
			if err != nil {
				return nil, nil, err
			}
		}
		if strings.Contains(v.Codecs, "hev1") {
			as = hevcMAs
		} else {
			as = mAs
		}
		video, err := as.AddNewRepresentationVideo(v.Bandwidth, v.Codecs, videoFormatsID2Name[v.ID], v.FrameRate, v.Width, v.Height)
		if err != nil {
			return nil, nil, err
		}
		video.Sar = &v.Sar
		err = video.AddNewBaseURL(v.BaseURL)
		if err != nil {
			return nil, nil, err
		}
		_, err = video.AddNewSegmentBase(v.SegmentBase.IndexRange, v.SegmentBase.Initialization)
		if err != nil {
			return nil, nil, err
		}
	}

	var ass []*mpd.AdaptationSet
	for _, a := range info.Data.Dash.Audio {
		if ass == nil {
			mAs, err = m.AddNewAdaptationSetAudio(a.MimeType, true, a.StartWithSap, "und")
			if err != nil {
				return nil, nil, err
			}
			hevcMAs, err = hevcM.AddNewAdaptationSetAudio(a.MimeType, true, a.StartWithSap, "und")
			if err != nil {
				return nil, nil, err
			}
			ass = append(ass, mAs, hevcMAs)
		}
		for _, as := range ass {
			audio, err := as.AddNewRepresentationAudio(44100, a.Bandwidth, a.Codecs, strconv.Itoa(a.ID))
			if err != nil {
				return nil, nil, err
			}
			audio.Sar = &a.Sar
			err = audio.AddNewBaseURL(a.BaseURL)
			if err != nil {
				return nil, nil, err
			}
			_, err = audio.AddNewSegmentBase(a.SegmentBase.IndexRange, a.SegmentBase.Initialization)
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return m, hevcM, nil
}

type Subtitle struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// https://api.bilibili.com/x/player/wbi/v2?bvid=BV1sw411Q7dr&cid=1248205613
func (c *Client) GetSubtitles(aid uint64, bvid string, cid uint64) ([]*Subtitle, error) {
	var url string
	if aid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/x/player/v2?aid=%d&cid=%d", aid, cid)
	} else if bvid != "" {
		url = fmt.Sprintf("https://api.bilibili.com/x/player/v2?bvid=%s&cid=%d", bvid, cid)
	} else {
		return nil, errors.New("aid and bvid are both empty")
	}
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := playerV2Info{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(info.Message)
	}
	r := make([]*Subtitle, len(info.Data.Subtitle.Subtitles))
	for i, s := range info.Data.Subtitle.Subtitles {
		r[i] = &Subtitle{
			Name: s.LanDoc,
			URL:  s.SubtitleURL,
		}
	}
	return r, nil
}

func (c *Client) ParsePGCPage(epid, seasonID uint64) (*VideoPageInfo, error) {
	var url string
	if epid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/pgc/view/web/season?ep_id=%d", epid)
	} else if seasonID != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/pgc/view/web/season?season_id=%d", seasonID)
	} else {
		return nil, errors.New("edId and season_id are both empty")
	}

	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := seasonInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(info.Message)
	}

	r := &VideoPageInfo{
		Title:      info.Result.Title,
		Actors:     info.Result.Actors,
		VideoInfos: make([]*VideoInfo, len(info.Result.Episodes)),
	}

	for i, v := range info.Result.Episodes {
		r.VideoInfos[i] = &VideoInfo{
			Bvid:       v.Bvid,
			Epid:       v.EpID,
			Cid:        v.Cid,
			Name:       v.ShareCopy,
			CoverImage: v.Cover,
		}
	}

	return r, nil
}

func (c *Client) GetPGCURL(epid, cid uint64, conf ...GetVideoURLConfig) (*VideoURL, error) {
	config := &GetVideoURLConf{
		Quality: Q1080PP,
	}
	for _, v := range conf {
		v(config)
	}
	config.fix()

	var url string
	if epid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/pgc/player/web/playurl?ep_id=%d&qn=%d&fourk=1&fnval=0", epid, config.Quality)
	} else if cid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/pgc/player/web/playurl?cid=%d&qn=%d&fourk=1&fnval=0", cid, config.Quality)
	} else {
		return nil, errors.New("edId and season_id are both empty")
	}

	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	info := pgcURLInfo{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, err
	}
	if info.Code != 0 {
		return nil, errors.New(info.Message)
	}

	return &VideoURL{
		AcceptDescription: info.Result.AcceptDescription,
		AcceptQuality:     info.Result.AcceptQuality,
		CurrentQuality:    info.Result.Quality,
		URL:               info.Result.Durl[0].URL,
	}, nil
}

func (c *Client) GetDashPGCURL(epid, cid uint64, conf ...GetDashVideoURLConfig) (*mpd.MPD, *mpd.MPD, error) {
	config := &GetDashVideoURLConf{}
	for _, v := range conf {
		v(config)
	}

	var (
		fnval    uint64 = 16
		extQuery string
	)

	if config.HDR {
		fnval |= 64
	}

	// 4K
	fnval |= 128
	// 8K
	fnval |= 1024

	var url string
	if epid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/pgc/player/web/playurl?ep_id=%d&fnval=%d%s", epid, fnval, extQuery)
	} else if cid != 0 {
		url = fmt.Sprintf("https://api.bilibili.com/pgc/player/web/playurl?cid=%d&fnval=%d%s", epid, fnval, extQuery)
	} else {
		return nil, nil, errors.New("edId and season_id are both empty")
	}
	req, err := c.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := c.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()
	info := dashPGCResp{}
	err = json.NewDecoder(resp.Body).Decode(&info)
	if err != nil {
		return nil, nil, err
	}
	if info.Code != 0 {
		return nil, nil, errors.New(info.Message)
	}
	videoFormatsID2Name := map[int]string{}
	for _, v := range info.Result.SupportFormats {
		videoFormatsID2Name[v.Quality] = v.NewDescription
	}
	m := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, fmt.Sprintf("PT%.2fS", info.Result.Dash.Duration), fmt.Sprintf("PT%.2fS", info.Result.Dash.MinBufferTime))
	hevcM := mpd.NewMPD(mpd.DASH_PROFILE_ONDEMAND, fmt.Sprintf("PT%.2fS", info.Result.Dash.Duration), fmt.Sprintf("PT%.2fS", info.Result.Dash.MinBufferTime))
	var as, mAs, hevcMAs *mpd.AdaptationSet
	for _, v := range info.Result.Dash.Video {
		if as == nil {
			mAs, err = m.AddNewAdaptationSetVideo(v.MimeType, "progressive", true, v.StartWithSap)
			if err != nil {
				return nil, nil, err
			}
			hevcMAs, err = hevcM.AddNewAdaptationSetVideo(v.MimeType, "progressive", true, v.StartWithSap)
			if err != nil {
				return nil, nil, err
			}
		}
		if strings.Contains(v.Codecs, "hev1") {
			as = hevcMAs
		} else {
			as = mAs
		}
		video, err := as.AddNewRepresentationVideo(v.Bandwidth, v.Codecs, videoFormatsID2Name[v.ID], v.FrameRate, v.Width, v.Height)
		if err != nil {
			return nil, nil, err
		}
		video.Sar = &v.Sar
		err = video.AddNewBaseURL(v.BaseURL)
		if err != nil {
			return nil, nil, err
		}
		_, err = video.AddNewSegmentBase(v.SegmentBase.IndexRange, v.SegmentBase.Initialization)
		if err != nil {
			return nil, nil, err
		}
	}

	var ass []*mpd.AdaptationSet
	for _, a := range info.Result.Dash.Audio {
		if ass == nil {
			mAs, err = m.AddNewAdaptationSetAudio(a.MimeType, true, a.StartWithSap, "und")
			if err != nil {
				return nil, nil, err
			}
			hevcMAs, err = hevcM.AddNewAdaptationSetAudio(a.MimeType, true, a.StartWithSap, "und")
			if err != nil {
				return nil, nil, err
			}
			ass = append(ass, mAs, hevcMAs)
		}
		for _, as := range ass {
			audio, err := as.AddNewRepresentationAudio(44100, a.Bandwidth, a.Codecs, strconv.Itoa(a.ID))
			if err != nil {
				return nil, nil, err
			}
			audio.Sar = &a.Sar
			err = audio.AddNewBaseURL(a.BaseURL)
			if err != nil {
				return nil, nil, err
			}
			_, err = audio.AddNewSegmentBase(a.SegmentBase.IndexRange, a.SegmentBase.Initialization)
			if err != nil {
				return nil, nil, err
			}
		}
	}
	return m, hevcM, nil
}
